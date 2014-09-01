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
	return 0xfd
}

func (field Field_STRING) Size() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
