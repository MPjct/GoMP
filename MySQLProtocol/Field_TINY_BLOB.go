package MySQLProtocol

type Field_TINY_BLOB struct {
	data string
}

func (field Field_TINY_BLOB) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_TINY_BLOB) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_TINY_BLOB) GetType() byte {
	return MYSQL_TYPE_TINY_BLOB
}

func (field Field_TINY_BLOB) Size() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
