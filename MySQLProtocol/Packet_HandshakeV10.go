package MySQLProtocol

import "math"

type Packet_HandshakeV10 struct {
	Packet

	protocol_version byte
	server_version   string
	connection_id    uint32
	auth_plugin_data string
	capability       uint32
	character_set    byte
	status           uint16
	auth_plugin_name string

	small bool
}

func (packet Packet_HandshakeV10) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += GetNulTerminatedStringSize(packet.server_version)
	size += 4
	size += 8
	size += 1
	size += 2
	if packet.small {
		return size
	}
	size += 1
	size += 2
	size += 2
	size += 1
	size += 10
	if Has_Flag(uint64(packet.capability), CLIENT_SECURE_CONNECTION) {
		size += uint64(math.Max(13, float64(len(packet.auth_plugin_data)-8)))
	}
	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH) {
		size += GetNulTerminatedStringSize(packet.auth_plugin_name)
	}

	return size
}

func (packet Packet_HandshakeV10) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, BuildFixedLengthInteger1(packet.protocol_version)...)
	data = append(data, BuildNulTerminatedString(packet.server_version)...)
	data = append(data, BuildFixedLengthInteger4(packet.connection_id)...)
	data = append(data, BuildFixedLengthString(packet.auth_plugin_data, 8)...)
	data = append(data, 0x00)
	data = append(data, BuildFixedLengthInteger2(uint16(packet.capability>>16))...)
	if packet.small {
		return data
	}

	data = append(data, BuildFixedLengthInteger1(packet.character_set)...)
	data = append(data, BuildFixedLengthInteger2(packet.status)...)
	data = append(data, BuildFixedLengthInteger2(uint16(packet.capability&0xFFFF))...)
	if len(packet.auth_plugin_data) > 8 {
		data = append(data, BuildFixedLengthInteger1(uint8(len(packet.auth_plugin_data)))...)
	} else {
		data = append(data, 0x00)
	}
	data = append(data, BuildFixedLengthString("", 10)...)
	if Has_Flag(uint64(packet.capability), CLIENT_SECURE_CONNECTION) {
		data = append(data, BuildFixedLengthString(packet.auth_plugin_data[8:], uint(math.Max(13, float64(len(packet.auth_plugin_data)-8))))...)
	}
	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH) {
		data = append(data, BuildNulTerminatedString(packet.auth_plugin_name)...)
	}

	return data
}

func (packet *Packet_HandshakeV10) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	packet.protocol_version = data.GetFixedLengthInteger1()
	packet.server_version = data.GetNulTerminatedString()
	packet.connection_id = data.GetFixedLengthInteger4()
	packet.auth_plugin_data = data.GetFixedLengthString(8)
	data.GetFixedLengthInteger1()
	packet.capability = uint32(data.GetFixedLengthInteger2()) << 16
	if !data.HasRemainingData() {
		packet.small = true
		return
	}
	packet.small = false
	packet.character_set = data.GetFixedLengthInteger1()
	packet.status = data.GetFixedLengthInteger2()
	packet.capability = uint32(Set_Flag(uint64(packet.capability), uint64(data.GetFixedLengthInteger2())))
	strlen := uint(data.GetFixedLengthInteger1())
	if strlen >= 8 {
		strlen -= 8
	}
	data.GetFixedLengthString(10)
	if !data.HasRemainingData() {
		return
	}
	if Has_Flag(uint64(packet.capability), CLIENT_SECURE_CONNECTION) {
		packet.auth_plugin_data += data.GetFixedLengthString(strlen)
	}
	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH) {
		packet.auth_plugin_name = data.GetNulTerminatedString()
	}
}
