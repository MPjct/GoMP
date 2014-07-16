package MySQLProtocol

type Packet_SSLRequest struct {
	Packet

	capability      uint32
	max_packet_size uint32
	username        string
	character_set   uint8
}

func (packet Packet_SSLRequest) GetPacketSize(context Context) (size uint64) {
	size += 4
	size += 4
	size += 1
	size += 23
	return size
}

func (packet Packet_SSLRequest) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	data = append(data, BuildFixedLengthInteger4(packet.capability)...)
	data = append(data, BuildFixedLengthInteger4(packet.max_packet_size)...)
	data = append(data, BuildFixedLengthInteger1(packet.character_set)...)
	data = append(data, BuildFixedLengthString("", 23)...)

	return data
}

func (packet *Packet_SSLRequest) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()

	packet.capability = data.GetFixedLengthInteger4()
	packet.max_packet_size = data.GetFixedLengthInteger4()
	packet.character_set = data.GetFixedLengthInteger1()
	data.GetFixedLengthString(23)
}
