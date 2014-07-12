package MySQLProtocol

type Packet_ERR struct {
    Packet
    err_code uint16
    sql_state string
    error_message string
}

func (packet Packet_ERR) GetPacketSize(context Context) (size uint64) {
    size += 1
    size += 2
    if Has_Flag(context.client_capability, CLIENT_PROTOCOL_41) {
        size += 1
        size += 5
    }
    size += uint64(len(packet.error_message))
    return size
}

func (packet Packet_ERR) ToPacket(context Context) (data []byte) {
    size := packet.GetPacketSize(context)
    
    data = make([]byte, 0, size+4)
    data = append(data, BuildFixedLengthInteger3(uint32(size))...)
    data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    data = append(data, 0xFF)
    data = append(data, BuildFixedLengthInteger2(packet.err_code)...)
    if Has_Flag(context.client_capability, CLIENT_PROTOCOL_41) {
        data = append(data, BuildFixedLengthString("#", 1)...)
        data = append(data, BuildFixedLengthString(packet.sql_state, 5)...)
    }
    data = append(data, BuildFixedLengthString(packet.error_message)...)
    return data
}

func (packet *Packet_ERR) FromPacket(context Context, data Proto) {
    data.GetFixedLengthInteger3()
    packet.sequence_id = data.GetFixedLengthInteger1()
    data.GetFixedLengthInteger1()
    packet.err_code = data.GetFixedLengthInteger2()
    if Has_Flag(context.client_capability, CLIENT_PROTOCOL_41) {
        data.GetFixedLengthInteger1()
        packet.sql_state = data.GetFixedLengthString(5)
    }
    packet.error_message = data.GetFixedLengthString()
}
