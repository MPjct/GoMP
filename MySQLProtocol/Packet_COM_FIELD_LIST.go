package MySQLProtocol

type Packet_COM_FIELD_LIST struct {
	Packet
    
    table string
    field string
}

func (packet Packet_COM_FIELD_LIST) GetPacketSize(context Context) (size uint64) {
	size += 1
    size += GetNulTerminatedStringSize(packet.table)
    size += GetFixedLengthStringSize(packet.field)
	return size
}

func (packet Packet_COM_FIELD_LIST) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_FIELD_LIST)
    data = append(data, BuildNulTerminatedString(packet.table)...)
    data = append(data, BuildFixedLengthString(packet.field)...)

	return data
}

func (packet *Packet_COM_FIELD_LIST) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
    packet.table = data.GetNulTerminatedString()
    packet.field = data.GetFixedLengthString()
}
