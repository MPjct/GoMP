package MySQLProtocol

type Packet_COM_SLEEP struct {
	Packet
}

func (packet Packet_COM_SLEEP) GetPacketSize(context Context) (size uint64) {
	size += 1
	return size
}

func (packet Packet_COM_SLEEP) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_SLEEP)

	return data
}

func (packet *Packet_COM_SLEEP) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
}
