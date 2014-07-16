package MySQLProtocol

type Packet_AuthMoreData struct {
	Packet

	auth_method_data string
}

func (packet Packet_AuthMoreData) GetPacketSize(context Context) (size uint64) {
    size += 1
    size += GetFixedLengthStringSize(packet.auth_method_data)
	return size
}

func (packet Packet_AuthMoreData) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
    data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    data = append(data, 0x01)
    data = append(data, BuildFixedLengthString(packet.auth_method_data)...)

	return data
}

func (packet *Packet_AuthMoreData) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
    data.GetFixedLengthInteger1()
    packet.auth_method_data = data.GetFixedLengthString()
}
