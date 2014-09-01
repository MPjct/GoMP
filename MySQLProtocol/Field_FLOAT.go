package MySQLProtocol

import "math"

type Field_FLOAT struct {
	data float32
}

func (field Field_FLOAT) Build() []byte {
	return BuildFixedLengthInteger4(math.Float32bits(field.data))
}

func (field *Field_FLOAT) Get(proto *Proto) {
	field.data = math.Float32frombits(proto.GetFixedLengthInteger4())
}

func (field Field_FLOAT) GetType() byte {
	return MYSQL_TYPE_FLOAT
}

func (field Field_FLOAT) Size() uint64 {
	return 4
}
