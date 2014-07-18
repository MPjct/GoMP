package MySQLProtocol

type Packet_ColumnDefinition41 struct {
	Packet
    
    catalog string
    schema string
    table string
    org_table string
    name string
    org_name string
    character_set uint16
    column_length uint32
    column_type uint8
    flags uint16
    decimals uint8
    default_values string
}

func (packet Packet_ColumnDefinition41) GetPacketSize(context Context) (size uint64) {
	size += GetLengthEncodedStringSize(packet.catalog)
    size += GetLengthEncodedStringSize(packet.schema)
    size += GetLengthEncodedStringSize(packet.table)
    size += GetLengthEncodedStringSize(packet.org_table)
    size += GetLengthEncodedStringSize(packet.name)
    size += GetLengthEncodedStringSize(packet.org_name)
    size += GetLengthEncodedIntegerSize(0x0C)
    size += 2
    size += 4
    size += 1
    size += 2
    size += 1
    size += 2
    if len(packet.default_values) > 0 {
        size += GetLengthEncodedStringSize(packet.default_values)
    }
	return size
}

func (packet Packet_ColumnDefinition41) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, BuildLengthEncodedString(packet.catalog)...)
    data = append(data, BuildLengthEncodedString(packet.schema)...)
    data = append(data, BuildLengthEncodedString(packet.table)...)
    data = append(data, BuildLengthEncodedString(packet.org_table)...)
    data = append(data, BuildLengthEncodedString(packet.name)...)
    data = append(data, BuildLengthEncodedString(packet.org_name)...)
    data = append(data, BuildLengthEncodedInteger(0x0C)...)
    data = append(data, BuildFixedLengthInteger2(packet.character_set)...)
    data = append(data, BuildFixedLengthInteger4(packet.column_length)...)
    data = append(data, BuildFixedLengthInteger1(packet.column_type)...)
    data = append(data, BuildFixedLengthInteger2(packet.flags)...)
    data = append(data, BuildFixedLengthInteger1(packet.decimals)...)
    data = append(data, BuildFixedLengthInteger2(0x0000)...)
    
    if len(packet.default_values) > 0 {
        data = append(data, BuildLengthEncodedString(packet.default_values)...)
    }
    
	return data
}

func (packet *Packet_ColumnDefinition41) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
    packet.catalog = data.GetLengthEncodedString()
    packet.schema = data.GetLengthEncodedString()
    packet.table = data.GetLengthEncodedString()
    packet.org_table = data.GetLengthEncodedString()
    packet.name = data.GetLengthEncodedString()
    packet.org_name = data.GetLengthEncodedString()
    data.GetLengthEncodedInteger()
    packet.character_set = data.GetFixedLengthInteger2()
    packet.column_length = data.GetFixedLengthInteger4()
    packet.column_type = data.GetFixedLengthInteger1()
    packet.flags = data.GetFixedLengthInteger2()
    packet.decimals = data.GetFixedLengthInteger1()
    data.GetFixedLengthInteger2()
    
    if data.HasRemainingData() {
        packet.default_values = data.GetLengthEncodedString()
    }
}
