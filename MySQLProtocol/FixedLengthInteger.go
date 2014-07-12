package MySQLProtocol

func BuildFixedLengthInteger1(value uint8) (data []byte) {
    data = make([]byte, 1)
	data[0] = byte(value >> 0 & 0xFF)
	return data
}

func (proto *Proto) GetFixedLengthInteger1() (value uint8) {
	value |= uint8(proto.data[proto.offset] & 0xFF)
	proto.offset += 1
	return value
}

func BuildFixedLengthInteger2(value uint16) (data []byte) {
    data = make([]byte, 2)
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	return data
}

func (proto *Proto) GetFixedLengthInteger2() (value uint16) {
	value |= uint16(proto.data[proto.offset+1]&0xFF) << 8
	value |= uint16(proto.data[proto.offset] & 0xFF)
	proto.offset += 2
	return value
}

func BuildFixedLengthInteger3(value uint32) (data []byte) {
    data = make([]byte, 3)
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	return data
}

func (proto *Proto) GetFixedLengthInteger3() (value uint32) {
	value |= uint32(proto.data[proto.offset+2]&0xFF) << 16
	value |= uint32(proto.data[proto.offset+1]&0xFF) << 8
	value |= uint32(proto.data[proto.offset] & 0xFF)
	proto.offset += 3
	return value
}

func BuildFixedLengthInteger4(value uint32) (data []byte) {
    data = make([]byte, 4)
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	data[3] = byte(value >> 24 & 0xFF)
	return data
}

func (proto *Proto) GetFixedLengthInteger4() (value uint32) {
	value |= uint32(proto.data[proto.offset+3]&0xFF) << 24
	value |= uint32(proto.data[proto.offset+2]&0xFF) << 16
	value |= uint32(proto.data[proto.offset+1]&0xFF) << 8
	value |= uint32(proto.data[proto.offset] & 0xFF)
	proto.offset += 4
	return value
}

func BuildFixedLengthInteger6(value uint64) (data []byte) {
    data = make([]byte, 6)
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	data[3] = byte(value >> 24 & 0xFF)
	data[4] = byte(value >> 32 & 0xFF)
	data[5] = byte(value >> 40 & 0xFF)
	return data
}

func (proto *Proto) GetFixedLengthInteger6() (value uint64) {
	value |= uint64(proto.data[proto.offset+5]&0xFF) << 40
	value |= uint64(proto.data[proto.offset+4]&0xFF) << 32
	value |= uint64(proto.data[proto.offset+3]&0xFF) << 24
	value |= uint64(proto.data[proto.offset+2]&0xFF) << 16
	value |= uint64(proto.data[proto.offset+1]&0xFF) << 8
	value |= uint64(proto.data[proto.offset] & 0xFF)
	proto.offset += 6
	return value
}

func BuildFixedLengthInteger8(value uint64) (data []byte) {
    data = make([]byte, 8)
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	data[3] = byte(value >> 24 & 0xFF)
	data[4] = byte(value >> 32 & 0xFF)
	data[5] = byte(value >> 40 & 0xFF)
	data[6] = byte(value >> 48 & 0xFF)
	data[7] = byte(value >> 56 & 0xFF)
	return data
}

func (proto *Proto) GetFixedLengthInteger8() (value uint64) {
	value |= uint64(proto.data[proto.offset+7]&0xFF) << 56
	value |= uint64(proto.data[proto.offset+6]&0xFF) << 48
	value |= uint64(proto.data[proto.offset+5]&0xFF) << 40
	value |= uint64(proto.data[proto.offset+4]&0xFF) << 32
	value |= uint64(proto.data[proto.offset+3]&0xFF) << 24
	value |= uint64(proto.data[proto.offset+2]&0xFF) << 16
	value |= uint64(proto.data[proto.offset+1]&0xFF) << 8
	value |= uint64(proto.data[proto.offset] & 0xFF)
	proto.offset += 8
	return value
}
