package MySQLProtocol

type Field_BIT struct {
	data string
}

func (field Field_BIT) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_BIT) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_BIT) GetType() byte {
	return MYSQL_TYPE_BIT
}

func (field Field_BIT) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
