package MySQLProtocol

type Field_BLOB struct {
	data string
}

func (field Field_BLOB) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_BLOB) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_BLOB) GetType() byte {
	return MYSQL_TYPE_BLOB
}

func (field Field_BLOB) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
