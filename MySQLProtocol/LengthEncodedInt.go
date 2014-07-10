package MySQLProtocol

func BuildLengthEncodedInt(value uint64) (data []byte) {
    if value < 251 {
        data = make([]byte, 1)
        data[0] = byte(value >> 0 & 0xFF)
        return data
    }
    
    if value < 0xFFFF {
        data = make([]byte, 3)
        data[0] = 0xFC
        data[1] = byte(value >> 0 & 0xFF)
        data[2] = byte(value >> 8 & 0xFF)
        return data
    }
    
    if value < 0xFFFFFF {
        data = make([]byte, 4)
        data[0] = 0xFD
        data[1] = byte(value >> 0 & 0xFF)
        data[2] = byte(value >> 8 & 0xFF)
        data[3] = byte(value >> 16 & 0xFF)
        return data
    }
    
    data = make([]byte, 9)
    data[0] = 0xFE
    data[1] = byte(value >> 0 & 0xFF)
    data[2] = byte(value >> 8 & 0xFF)
    data[3] = byte(value >> 16 & 0xFF)
    data[4] = byte(value >> 24 & 0xFF)
    data[5] = byte(value >> 32 & 0xFF)
    data[6] = byte(value >> 40 & 0xFF)
    data[7] = byte(value >> 48 & 0xFF)
    data[8] = byte(value >> 56 & 0xFF)
    return data
}
