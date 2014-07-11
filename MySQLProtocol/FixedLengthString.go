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

func (packet Packet) GetFixedLengthString(size ...uint) (value string) {
	var datasize uint
	if len(size) == 0 {
		datasize = uint(len(packet.data)) - packet.offset
	} else {
		datasize = size[0]
	}
	value = string(packet.data[packet.offset : packet.offset+datasize])
	packet.offset += datasize
	return value
}
