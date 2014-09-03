package MySQLProtocol

type Packet_COM_STMT_FETCH struct {
	Packet

	statement_id uint32
    num_rows uint32
}

func (packet Packet_COM_STMT_FETCH) GetPacketSize(context Context) (size uint64) {
    size += 1
	size += 4
    size += 4
	return size
}

func (packet Packet_COM_STMT_FETCH) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	data = append(data, COM_STMT_FETCH)
	data = append(data, BuildFixedLengthInteger4(packet.statement_id)...)
    data = append(data, BuildFixedLengthInteger4(packet.num_rows)...)

	return data
}

func (packet *Packet_COM_STMT_FETCH) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	packet.statement_id = data.GetFixedLengthInteger4()
    packet.num_rows = data.GetFixedLengthInteger4()
}
