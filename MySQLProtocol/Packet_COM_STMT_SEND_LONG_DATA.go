package MySQLProtocol

type Packet_COM_STMT_SEND_LONG_DATA struct {
	Packet

    statement_id uint32
    param_id uint16
    data string
}

func (packet Packet_COM_STMT_SEND_LONG_DATA) GetPacketSize(context Context) (size uint64) {
    size += 1
    size += 4
    size += 2
    size += GetFixedLengthStringSize(packet.data)
	return size
}

func (packet Packet_COM_STMT_SEND_LONG_DATA) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    
    data = append(data, COM_STMT_SEND_LONG_DATA)
    
    data = append(data, BuildFixedLengthInteger4(packet.statement_id)...)
    data = append(data, BuildFixedLengthInteger2(packet.param_id)...)
    data = append(data, BuildFixedLengthString(packet.data)...)

	return data
}

func (packet *Packet_COM_STMT_SEND_LONG_DATA) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
    data.GetFixedLengthInteger1()
	packet.statement_id = data.GetFixedLengthInteger4()
	packet.param_id = data.GetFixedLengthInteger2()
	packet.data = data.GetFixedLengthString()
}
