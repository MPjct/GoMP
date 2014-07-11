package MySQLProtocol

func BuildFixedLengthString(value string, size ...uint) (data []byte) {
	var datasize uint
	if len(size) == 0 {
		datasize = uint(len(value))
		data = make([]byte, datasize)
	} else if size[0] > uint(len(value)) {
		data = make([]byte, size[0])
		datasize = uint(len(value))
	} else {
		data = make([]byte, size[0])
		datasize = size[0]
	}
	copy(data, []byte(value[:datasize]))
	return data
}

func (proto Proto) GetFixedLengthString(size ...uint) (value string) {
	var datasize uint
	if len(size) == 0 {
		datasize = uint(len(proto.data)) - proto.offset
	} else {
		datasize = size[0]
	}
	value = string(proto.data[proto.offset : proto.offset+datasize])
	proto.offset += datasize
	return value
}
