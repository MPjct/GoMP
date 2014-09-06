package MySQLProtocol

type Field_INT24 struct {
	data uint32
}

func (field Field_INT24) Build() []byte {
	return BuildFixedLengthInteger4(field.data)
}

func (field *Field_INT24) Get(proto *Proto) {
	field.data = proto.GetFixedLengthInteger4()
}

func (field Field_INT24) GetType() byte {
	return MYSQL_TYPE_INT24
}

func (field Field_INT24) PacketSize() uint64 {
	return 4
}
