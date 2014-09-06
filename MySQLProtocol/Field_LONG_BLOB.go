package MySQLProtocol

type Field_LONG_BLOB struct {
	data string
}

func (field Field_LONG_BLOB) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_LONG_BLOB) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_LONG_BLOB) GetType() byte {
	return MYSQL_TYPE_LONG_BLOB
}

func (field Field_LONG_BLOB) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
