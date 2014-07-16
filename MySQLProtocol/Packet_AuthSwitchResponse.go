package MySQLProtocol

type Packet_AuthSwitchResponse struct {
	Packet

	auth_plugin_response string
}

func (packet Packet_AuthSwitchResponse) GetPacketSize(context Context) (size uint64) {
    size += GetFixedLengthStringSize(packet.auth_plugin_response)
	return size
}

func (packet Packet_AuthSwitchResponse) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
    data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    data = append(data, BuildFixedLengthString(packet.auth_plugin_response)...)

	return data
}

func (packet *Packet_AuthSwitchResponse) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
    packet.auth_plugin_response = data.GetFixedLengthString()
}
