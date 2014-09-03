package MySQLProtocol

type Packet_COM_SET_OPTION struct {
	Packet

	option uint16
}

func (packet Packet_COM_SET_OPTION) GetPacketSize(context Context) (size uint64) {
    size += 1
	size += 2
	return size
}

func (packet Packet_COM_SET_OPTION) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	data = append(data, COM_SET_OPTION)
	data = append(data, BuildFixedLengthInteger2(packet.option)...)

	return data
}

func (packet *Packet_COM_SET_OPTION) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	packet.option = data.GetFixedLengthInteger2()
}
