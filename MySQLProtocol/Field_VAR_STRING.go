package MySQLProtocol

type Field_VAR_STRING struct {
	data string
}

func (field Field_VAR_STRING) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_VAR_STRING) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_VAR_STRING) GetType() byte {
	return MYSQL_TYPE_VAR_STRING
}

func (field Field_VAR_STRING) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
