package MySQLProtocol

type Packet_HandshakeV9 struct {
	Packet

	protocol_version byte
	server_version   string
	connection_id    uint32
	auth_plugin_data string
}

func (packet Packet_HandshakeV9) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += GetNulTerminatedStringSize(packet.server_version)
	size += 4
    size += GetNulTerminatedStringSize(packet.auth_plugin_data)

	return size
}

func (packet Packet_HandshakeV9) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, BuildFixedLengthInteger1(packet.protocol_version)...)
	data = append(data, BuildNulTerminatedString(packet.server_version)...)
	data = append(data, BuildFixedLengthInteger4(packet.connection_id)...)
    data = append(data, BuildNulTerminatedString(packet.auth_plugin_data)...)

	return data
}

func (packet *Packet_HandshakeV9) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	packet.protocol_version = data.GetFixedLengthInteger1()
	packet.server_version = data.GetNulTerminatedString()
	packet.connection_id = data.GetFixedLengthInteger4()
	packet.auth_plugin_data = data.GetNulTerminatedString()
}
