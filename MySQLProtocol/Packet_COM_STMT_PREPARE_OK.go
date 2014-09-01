package MySQLProtocol

type Packet_COM_STMT_PREPARE_OK struct {
	Packet

	status        uint8
	statement_id  uint32
	num_columns   uint16
	num_params    uint16
	warning_count uint16
}

func (packet Packet_COM_STMT_PREPARE_OK) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += 4
	size += 2
	size += 2
	size += 1
	size += 2
	return size
}

func (packet Packet_COM_STMT_PREPARE_OK) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	data = append(data, BuildFixedLengthInteger1(packet.status)...)
	data = append(data, BuildFixedLengthInteger4(packet.statement_id)...)
	data = append(data, BuildFixedLengthInteger2(packet.num_columns)...)
	data = append(data, BuildFixedLengthInteger2(packet.num_params)...)
	data = append(data, BuildFixedLengthInteger1(0x00)...)
	data = append(data, BuildFixedLengthInteger2(packet.warning_count)...)

	return data
}

func (packet *Packet_COM_STMT_PREPARE_OK) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	packet.status = data.GetFixedLengthInteger1()
	packet.statement_id = data.GetFixedLengthInteger4()
	packet.num_columns = data.GetFixedLengthInteger2()
	packet.num_params = data.GetFixedLengthInteger2()
	data.GetFixedLengthInteger1()
	packet.warning_count = data.GetFixedLengthInteger2()
}
