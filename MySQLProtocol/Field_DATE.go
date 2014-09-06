package MySQLProtocol

type Field_DATE struct {
	year  uint16
	month uint8
	day   uint8
}

func (field Field_DATE) Build() []byte {
	length := field.Size()
	data := make([]byte, 0, length)

	data = append(data, byte(length))

	if length >= 4 {
		data = append(data, BuildFixedLengthInteger2(field.year)...)
		data = append(data, BuildFixedLengthInteger1(field.month)...)
		data = append(data, BuildFixedLengthInteger1(field.day)...)
	}

	return data
}

func (field *Field_DATE) Get(proto *Proto) {
	length := proto.GetFixedLengthInteger1()

	if length >= 4 {
		field.year = proto.GetFixedLengthInteger2()
		field.month = proto.GetFixedLengthInteger1()
		field.day = proto.GetFixedLengthInteger1()
	}
}

func (field Field_DATE) GetType() byte {
	return MYSQL_TYPE_DATE
}

func (field Field_DATE) Size() uint64 {
	if field.year != 0 || field.month != 0 || field.day != 0 {
		return 4
	}

	return 0
}

func (field Field_DATE) PacketSize() uint64 {
    return field.Size()+1
}
