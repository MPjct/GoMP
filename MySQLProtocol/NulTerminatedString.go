package MySQLProtocol

func BuildNulTerminatedString(value string) (data []byte) {
    data = make([]byte, len(value)+1)
    copy(data, []byte(value))
    return data
}

func (packet Packet) GetNulTerminatedString() (value string) {
    var strlen uint
    for packet.data[packet.offset + strlen] != 0x00 {
        strlen++
    }
    value = string(packet.data[packet.offset: packet.offset + strlen])
    packet.offset += strlen+1
    return value
}
