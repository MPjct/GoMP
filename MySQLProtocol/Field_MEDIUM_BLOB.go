package MySQLProtocol

type Field_MEDIUM_BLOB struct {
	data string
}

func (field Field_MEDIUM_BLOB) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_MEDIUM_BLOB) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_MEDIUM_BLOB) GetType() byte {
	return MYSQL_TYPE_MEDIUM_BLOB
}

func (field Field_MEDIUM_BLOB) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
