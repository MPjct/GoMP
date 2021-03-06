package MySQLProtocol

type Packet_COM_QUERY struct {
	Packet

	query string
}

func (packet Packet_COM_QUERY) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += GetFixedLengthStringSize(packet.query)
	return size
}

func (packet Packet_COM_QUERY) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_QUERY)
	data = append(data, BuildFixedLengthString(packet.query)...)

	return data
}

func (packet *Packet_COM_QUERY) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	packet.query = data.GetFixedLengthString()
}
