package MySQLProtocol

type Packet_LOCAL_INFILE_Request struct {
	Packet
    
    filename string
}

func (packet Packet_LOCAL_INFILE_Request) GetPacketSize(context Context) (size uint64) {
	size += 1
    size += GetFixedLengthStringSize(packet.filename)
	return size
}

func (packet Packet_LOCAL_INFILE_Request) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, LOCAL_INFILE)
    data = append(data, BuildFixedLengthString(packet.filename)...)
	return data
}

func (packet *Packet_LOCAL_INFILE_Request) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
    packet.filename = data.GetFixedLengthString()
}
