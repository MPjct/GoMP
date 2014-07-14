package MySQLProtocol

func BuildLengthEncodedString(value string) (data []byte) {
	strlen := BuildLengthEncodedInteger(uint64(len(value)))
	data = make([]byte, 0, len(strlen)+len(value))
	data = append(data, strlen...)
	data = append(data, []byte(value)...)
	return data
}

func (proto *Proto) GetLengthEncodedString() (value string) {
	strlen := uint(proto.GetLengthEncodedInteger())
	return proto.GetFixedLengthString(strlen)
}

func GetLengthEncodedStringSize(value string) uint64 {
	return GetLengthEncodedIntegerSize(uint64(len(value))) + uint64(len(value))
}
