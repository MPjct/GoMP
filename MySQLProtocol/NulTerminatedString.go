package MySQLProtocol

func BuildNulTerminatedString(value string) (data []byte) {
	data = make([]byte, len(value)+1)
	copy(data, []byte(value))
	return data
}

func (proto *Proto) GetNulTerminatedString() (value string) {
	var strlen uint
	for proto.data[proto.offset+strlen] != 0x00 {
		strlen++
	}
	value = string(proto.data[proto.offset:proto.offset+strlen])
	proto.offset += strlen + 1
	return value
}

func GetNulTerminatedStringSize(value string) uint64 {
    return uint64(len(value)) + 1
}
