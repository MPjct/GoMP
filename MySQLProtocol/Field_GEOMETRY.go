package MySQLProtocol

type Field_GEOMETRY struct {
	data string
}

func (field Field_GEOMETRY) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_GEOMETRY) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_GEOMETRY) GetType() byte {
	return MYSQL_TYPE_GEOMETRY
}

func (field Field_GEOMETRY) Size() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
