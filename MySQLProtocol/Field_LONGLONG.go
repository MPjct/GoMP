package MySQLProtocol

type Field_LONGLONG struct {
	data uint64
}

func (field Field_LONGLONG) Build() []byte {
	return BuildFixedLengthInteger8(field.data)
}

func (field *Field_LONGLONG) Get(proto *Proto) {
	field.data = proto.GetFixedLengthInteger8()
}

func (field Field_LONGLONG) GetType() byte {
	return MYSQL_TYPE_LONGLONG
}

func (field Field_LONGLONG) PacketSize() uint64 {
	return 8
}
