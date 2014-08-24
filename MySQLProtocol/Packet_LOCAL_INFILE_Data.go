package MySQLProtocol

type Packet_LOCAL_INFILE_Data struct {
	Packet

	data string
}

func (packet Packet_LOCAL_INFILE_Data) GetPacketSize(context Context) (size uint64) {
	size += GetFixedLengthStringSize(packet.data)
	return size
}

func (packet Packet_LOCAL_INFILE_Data) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, BuildFixedLengthString(packet.data)...)
	return data
}

func (packet *Packet_LOCAL_INFILE_Data) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	packet.data = data.GetFixedLengthString()
}
