package MySQLProtocol

type Field_NULL struct {
}

func (field Field_NULL) Build() []byte {
	return []byte{}
}

func (field *Field_NULL) Get(proto *Proto) {
}

func (field Field_NULL) GetType() byte {
	return MYSQL_TYPE_NULL
}

func (field Field_NULL) Size() uint64 {
	return 0
}
