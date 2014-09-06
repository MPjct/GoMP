package MySQLProtocol

type Field_LONG struct {
	data uint32
}

func (field Field_LONG) Build() []byte {
	return BuildFixedLengthInteger4(field.data)
}

func (field *Field_LONG) Get(proto *Proto) {
	field.data = proto.GetFixedLengthInteger4()
}

func (field Field_LONG) GetType() byte {
	return MYSQL_TYPE_LONG
}

func (field Field_LONG) PacketSize() uint64 {
	return 4
}
