package MySQLProtocol

type Field_YEAR struct {
	data uint16
}

func (field Field_YEAR) Build() []byte {
	return BuildFixedLengthInteger2(field.data)
}

func (field *Field_YEAR) Get(proto *Proto) {
	field.data = proto.GetFixedLengthInteger2()
}

func (field Field_YEAR) GetType() byte {
	return MYSQL_TYPE_YEAR
}

func (field Field_YEAR) Size() uint64 {
	return 2
}
