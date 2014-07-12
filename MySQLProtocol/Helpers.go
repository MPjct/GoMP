package MySQLProtocol

func Has_Flag(value uint64, flag uint64) (result bool) {
    return (value & flag) == flag
}

func Set_Flag(value uint64, flag uint64) (result uint64) {
    return value|flag
}

func Remove_Flag(value uint64, flag uint64) (result uint64) {
    return value&flag
}

func Toggle_Flag(value uint64, flag uint64) (result uint64) {
    return value^flag
}
