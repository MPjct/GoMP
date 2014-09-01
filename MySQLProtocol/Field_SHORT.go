package MySQLProtocol

type Field_SHORT struct {
	data uint16
}

func (field Field_SHORT) Build() []byte {
	return BuildFixedLengthInteger2(field.data)
}

func (field *Field_SHORT) Get(proto *Proto) {
	field.data = proto.GetFixedLengthInteger2()
}

func (field Field_SHORT) GetType() byte {
	return MYSQL_TYPE_SHORT
}

func (field Field_SHORT) Size() uint64 {
	return 2
}
