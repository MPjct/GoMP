package MySQLProtocol

type Packet_COM_REFRESH struct {
	Packet
	sub_command uint8
}

func (packet Packet_COM_REFRESH) GetPacketSize(context Context) (size uint64) {
	size += 2
	return size
}

func (packet Packet_COM_REFRESH) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_REFRESH)
	data = append(data, packet.sub_command)
	return data
}

func (packet *Packet_COM_REFRESH) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	packet.sub_command = data.GetFixedLengthInteger1()
}
