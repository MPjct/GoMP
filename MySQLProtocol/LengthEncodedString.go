package MySQLProtocol

func BuildLengthEncodedString(value string) (data []byte) {
	strlen := BuildLengthEncodedInteger(uint64(len(value)))
	data = make([]byte, len(strlen)+len(value))
	copy(data, strlen)
	copy(data[len(strlen):], []byte(value))
	return data
}
