package MySQLProtocol

type Field_VARCHAR struct {
	data string
}

func (field Field_VARCHAR) Build() []byte {
	return BuildLengthEncodedString(field.data)
}

func (field *Field_VARCHAR) Get(proto *Proto) {
	field.data = proto.GetLengthEncodedString()
}

func (field Field_VARCHAR) GetType() byte {
	return MYSQL_TYPE_VARCHAR
}

func (field Field_VARCHAR) PacketSize() uint64 {
	return GetLengthEncodedStringSize(field.data)
}
