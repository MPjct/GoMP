package MySQLProtocol

type Packet_ColumnDefinition320 struct {
	Packet

	table          string
	name           string
	column_length  uint32
	column_type    uint8
	flags          uint16
	decimals       uint8
	default_values string
}

func (packet Packet_ColumnDefinition320) GetPacketSize(context Context) (size uint64) {
	size += GetLengthEncodedStringSize(packet.table)
	size += GetLengthEncodedStringSize(packet.name)
	size += GetLengthEncodedIntegerSize(0x03)
	size += 3
	size += GetLengthEncodedIntegerSize(0x01)
	size += 1

	if Has_Flag(context.capability, CLIENT_LONG_FLAG) {
		size += GetLengthEncodedIntegerSize(0x03)
		size += 2
		size += 1
	} else {
		size += GetLengthEncodedIntegerSize(0x02)
		size += 1
		size += 1
	}

	if len(packet.default_values) > 0 {
		size += GetLengthEncodedStringSize(packet.default_values)
	}
	return size
}

func (packet Packet_ColumnDefinition320) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, BuildLengthEncodedString(packet.table)...)
	data = append(data, BuildLengthEncodedString(packet.name)...)
	data = append(data, BuildLengthEncodedInteger(0x03)...)
	data = append(data, BuildFixedLengthInteger3(packet.column_length)...)
	data = append(data, BuildLengthEncodedInteger(0x01)...)
	data = append(data, BuildFixedLengthInteger1(packet.column_type)...)
	if Has_Flag(context.capability, CLIENT_LONG_FLAG) {
		data = append(data, BuildLengthEncodedInteger(0x03)...)
		data = append(data, BuildFixedLengthInteger2(packet.flags)...)
		data = append(data, BuildFixedLengthInteger1(packet.decimals)...)
	} else {
		data = append(data, BuildLengthEncodedInteger(0x02)...)
		data = append(data, BuildFixedLengthInteger1(uint8(packet.flags))...)
		data = append(data, BuildFixedLengthInteger1(packet.decimals)...)
	}

	if len(packet.default_values) > 0 {
		data = append(data, BuildLengthEncodedString(packet.default_values)...)
	}

	return data
}

func (packet *Packet_ColumnDefinition320) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	packet.table = data.GetLengthEncodedString()
	packet.name = data.GetLengthEncodedString()
	data.GetLengthEncodedInteger()
	packet.column_length = data.GetFixedLengthInteger3()
	data.GetLengthEncodedInteger()
	packet.column_type = data.GetFixedLengthInteger1()
	data.GetLengthEncodedInteger()
	if Has_Flag(context.capability, CLIENT_LONG_FLAG) {
		packet.flags = data.GetFixedLengthInteger2()
		packet.decimals = data.GetFixedLengthInteger1()
	} else {
		packet.flags = uint16(data.GetFixedLengthInteger1())
		packet.decimals = data.GetFixedLengthInteger1()
	}

	if data.HasRemainingData() {
		packet.default_values = data.GetLengthEncodedString()
	}
}
