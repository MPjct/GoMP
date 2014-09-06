package MySQLProtocol

type Field_DATETIME struct {
	year         uint16
	month        uint8
	day          uint8
	hour         uint8
	minute       uint8
	second       uint8
	micro_second uint32
}

func (field Field_DATETIME) Build() []byte {
	length := field.Size()
	data := make([]byte, 0, length)

	data = append(data, byte(length))

	if length >= 4 {
		data = append(data, BuildFixedLengthInteger2(field.year)...)
		data = append(data, BuildFixedLengthInteger1(field.month)...)
		data = append(data, BuildFixedLengthInteger1(field.day)...)
	}

	if length >= 7 {
		data = append(data, BuildFixedLengthInteger1(field.hour)...)
		data = append(data, BuildFixedLengthInteger1(field.minute)...)
		data = append(data, BuildFixedLengthInteger1(field.second)...)
	}

	if length == 11 {
		data = append(data, BuildFixedLengthInteger4(field.micro_second)...)
	}

	return data
}

func (field *Field_DATETIME) Get(proto *Proto) {
	length := proto.GetFixedLengthInteger1()

	if length >= 4 {
		field.year = proto.GetFixedLengthInteger2()
		field.month = proto.GetFixedLengthInteger1()
		field.day = proto.GetFixedLengthInteger1()
	}

	if length >= 7 {
		field.hour = proto.GetFixedLengthInteger1()
		field.minute = proto.GetFixedLengthInteger1()
		field.second = proto.GetFixedLengthInteger1()
	}

	if length == 11 {
		field.micro_second = proto.GetFixedLengthInteger4()
	}
}

func (field Field_DATETIME) GetType() byte {
	return MYSQL_TYPE_DATETIME
}

func (field Field_DATETIME) Size() uint64 {
	if field.micro_second != 0 {
		return 11
	}

	if field.hour != 0 || field.minute != 0 || field.second != 0 {
		return 7
	}

	if field.year != 0 || field.month != 0 || field.day != 0 {
		return 4
	}

	return 0
}

func (field Field_DATETIME) PacketSize() uint64 {
	return field.Size() + 1
}
