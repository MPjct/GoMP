package MySQLProtocol

type Field_NEWDECIMAL struct {
	data string
}

func (field Field_NEWDECIMAL) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_NEWDECIMAL) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_NEWDECIMAL) GetType() byte {
	return MYSQL_TYPE_NEWDECIMAL
}

func (field Field_NEWDECIMAL) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
