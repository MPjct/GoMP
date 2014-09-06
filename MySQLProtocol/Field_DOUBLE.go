package MySQLProtocol

import "math"

type Field_DOUBLE struct {
	data float64
}

func (field Field_DOUBLE) Build() []byte {
	return BuildFixedLengthInteger8(math.Float64bits(field.data))
}

func (field *Field_DOUBLE) Get(proto *Proto) {
	field.data = math.Float64frombits(proto.GetFixedLengthInteger8())
}

func (field Field_DOUBLE) GetType() byte {
	return MYSQL_TYPE_DOUBLE
}

func (field Field_DOUBLE) PacketSize() uint64 {
	return 8
}
