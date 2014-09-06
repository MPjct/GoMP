package MySQLProtocol

type Field_DECIMAL struct {
	data string
}

func (field Field_DECIMAL) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_DECIMAL) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_DECIMAL) GetType() byte {
	return MYSQL_TYPE_DECIMAL
}

func (field Field_DECIMAL) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
