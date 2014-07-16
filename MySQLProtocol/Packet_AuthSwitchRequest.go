package MySQLProtocol

type Packet_AuthSwitchRequest struct {
	Packet

	protocol_version byte
	plugin_name      string
	auth_plugin_data string
}

func (packet Packet_AuthSwitchRequest) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += GetNulTerminatedStringSize(packet.plugin_name)
	size += GetFixedLengthStringSize(packet.auth_plugin_data)

	return size
}

func (packet Packet_AuthSwitchRequest) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, BuildFixedLengthInteger1(0xFE)...)
	data = append(data, BuildNulTerminatedString(packet.plugin_name)...)
	data = append(data, BuildFixedLengthString(packet.auth_plugin_data)...)

	return data
}

func (packet *Packet_AuthSwitchRequest) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	packet.plugin_name = data.GetNulTerminatedString()
	packet.auth_plugin_data = data.GetFixedLengthString()
}
