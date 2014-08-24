package MySQLProtocol

type Packet_COM_PROCESS_INFO struct {
	Packet
}

func (packet Packet_COM_PROCESS_INFO) GetPacketSize(context Context) (size uint64) {
	size += 1
	return size
}

func (packet Packet_COM_PROCESS_INFO) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_PROCESS_INFO)

	return data
}

func (packet *Packet_COM_PROCESS_INFO) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
}
