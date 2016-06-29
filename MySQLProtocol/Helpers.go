package MySQLProtocol

func Has_Flag(value uint64, flag uint64) bool {
	return (value & flag) == flag
}

func Has_Flag_uint16(value uint16, flag uint16) bool {
	return (value & flag) == flag
}

func Has_Flag_uint64(value uint64, flag uint64) bool {
	return (value & flag) == flag
}

func Set_Flag(value uint64, flag uint64) uint64 {
	return value | flag
}

func Remove_Flag(value uint64, flag uint64) uint64 {
	return value &^ flag
}
