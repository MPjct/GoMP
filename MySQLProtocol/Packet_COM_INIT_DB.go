package MySQLProtocol

type Packet_COM_INIT_DB struct {
	Packet

	schema string
}

func (packet Packet_COM_INIT_DB) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += GetFixedLengthStringSize(packet.schema)
	return size
}

func (packet Packet_COM_INIT_DB) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_INIT_DB)
	data = append(data, BuildFixedLengthString(packet.schema)...)

	return data
}

func (packet *Packet_COM_INIT_DB) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	packet.schema = data.GetFixedLengthString()
}
