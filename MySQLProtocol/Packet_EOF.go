package MySQLProtocol

type Packet_EOF struct {
    Packet
    warning_count uint16
    status_flags uint16
}

func (packet Packet_EOF) GetPacketSize(context Context) (size uint64) {
    size += 1
    if Has_Flag(context.capability, CLIENT_PROTOCOL_41) {
        size += 2
        size += 2
    }
    return size
}

func (packet Packet_EOF) ToPacket(context Context) (data []byte) {
    size := packet.GetPacketSize(context)
    
    data = make([]byte, 0, size+4)
    data = append(data, BuildFixedLengthInteger3(uint32(size))...)
    data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    data = append(data, 0xFE)
    if Has_Flag(context.capability, CLIENT_PROTOCOL_41) {
        data = append(data, BuildFixedLengthInteger2(packet.warning_count)...)
        data = append(data, BuildFixedLengthInteger2(packet.status_flags)...)
    }
    return data
}

func (packet *Packet_EOF) FromPacket(context Context, data Proto) {
    data.GetFixedLengthInteger3()
    packet.sequence_id = data.GetFixedLengthInteger1()
    data.GetFixedLengthInteger1()
    if Has_Flag(context.capability, CLIENT_PROTOCOL_41) {
        packet.warning_count = data.GetFixedLengthInteger2()
        packet.status_flags = data.GetFixedLengthInteger2()
    }
}
