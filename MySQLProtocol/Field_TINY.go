package MySQLProtocol

type Field_TINY struct {
	data uint8
}

func (field Field_TINY) Build() []byte {
	return BuildFixedLengthInteger1(field.data)
}

func (field *Field_TINY) Get(proto *Proto) {
	field.data = proto.GetFixedLengthInteger1()
}

func (field Field_TINY) GetType() byte {
	return MYSQL_TYPE_TINY
}

func (field Field_TINY) PacketSize() uint64 {
	return 1
}
