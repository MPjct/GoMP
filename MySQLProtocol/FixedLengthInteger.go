package MySQLProtocol

func BuildFixedLengthInteger1(value uint8) (data [1]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	return data
}

func (packet Packet) GetFixedLengthInteger1() (value uint8) {
    value |= uint8(packet.data[packet.offset] & 0xFF)
    packet.offset += 1
    return value
}

func BuildFixedLengthInteger2(value uint16) (data [2]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	return data
}

func (packet Packet) GetFixedLengthInteger2() (value uint16) {
    value |= uint16(packet.data[packet.offset+1] & 0xFF) << 8
    value |= uint16(packet.data[packet.offset] & 0xFF)
    packet.offset += 2
    return value
}

func BuildFixedLengthInteger3(value uint32) (data [3]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	return data
}

func (packet Packet) GetFixedLengthInteger3() (value uint32) {
    value |= uint32(packet.data[packet.offset+2] & 0xFF) << 16
    value |= uint32(packet.data[packet.offset+1] & 0xFF) << 8
    value |= uint32(packet.data[packet.offset] & 0xFF)
    packet.offset += 3
    return value
}

func BuildFixedLengthInteger4(value uint32) (data [4]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	data[3] = byte(value >> 24 & 0xFF)
	return data
}

func (packet Packet) GetFixedLengthInteger4() (value uint32) {
    value |= uint32(packet.data[packet.offset+3] & 0xFF) << 24
    value |= uint32(packet.data[packet.offset+2] & 0xFF) << 16
    value |= uint32(packet.data[packet.offset+1] & 0xFF) << 8
    value |= uint32(packet.data[packet.offset] & 0xFF)
    packet.offset += 4
    return value
}

func BuildFixedLengthInteger6(value uint64) (data [6]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	data[3] = byte(value >> 24 & 0xFF)
	data[4] = byte(value >> 32 & 0xFF)
	data[5] = byte(value >> 40 & 0xFF)
	return data
}

func (packet Packet) GetFixedLengthInteger6() (value uint64) {
    value |= uint64(packet.data[packet.offset+5] & 0xFF) << 40
    value |= uint64(packet.data[packet.offset+4] & 0xFF) << 32
    value |= uint64(packet.data[packet.offset+3] & 0xFF) << 24
    value |= uint64(packet.data[packet.offset+2] & 0xFF) << 16
    value |= uint64(packet.data[packet.offset+1] & 0xFF) << 8
    value |= uint64(packet.data[packet.offset] & 0xFF)
    packet.offset += 6
    return value
}

func BuildFixedLengthInteger8(value uint64) (data [8]byte) {
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

func (packet Packet) GetFixedLengthInteger8() (value uint64) {
    value |= uint64(packet.data[packet.offset+7] & 0xFF) << 56
    value |= uint64(packet.data[packet.offset+6] & 0xFF) << 48
    value |= uint64(packet.data[packet.offset+5] & 0xFF) << 40
    value |= uint64(packet.data[packet.offset+4] & 0xFF) << 32
    value |= uint64(packet.data[packet.offset+3] & 0xFF) << 24
    value |= uint64(packet.data[packet.offset+2] & 0xFF) << 16
    value |= uint64(packet.data[packet.offset+1] & 0xFF) << 8
    value |= uint64(packet.data[packet.offset] & 0xFF)
    packet.offset += 8
    return value
}
