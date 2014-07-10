package MySQLProtocol

func BuildFixedInt1(value uint8) (data [1]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	return data
}

func BuildFixedInt2(value uint16) (data [2]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	return data
}

func BuildFixedInt3(value uint32) (data [3]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	return data
}

func BuildFixedInt4(value uint32) (data [4]byte) {
	data[0] = byte(value >> 0 & 0xFF)
	data[1] = byte(value >> 8 & 0xFF)
	data[2] = byte(value >> 16 & 0xFF)
	data[3] = byte(value >> 24 & 0xFF)
	return data
}

func BuildFixedInt8(value uint64) (data [8]byte) {
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
