package MySQLProtocol

type Field_SET struct {
	data string
}

func (field Field_SET) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_SET) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_SET) GetType() byte {
	return MYSQL_TYPE_SET
}

func (field Field_SET) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
