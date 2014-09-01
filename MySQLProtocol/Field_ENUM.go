package MySQLProtocol

type Field_ENUM struct {
	data string
}

func (field Field_ENUM) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_ENUM) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_ENUM) GetType() byte {
	return MYSQL_TYPE_ENUM
}

func (field Field_ENUM) Size() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
