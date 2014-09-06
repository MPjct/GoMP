package MySQLProtocol

type Field_STRING struct {
	data string
}

func (field Field_STRING) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_STRING) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_STRING) GetType() byte {
	return MYSQL_TYPE_STRING
}

func (field Field_STRING) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
