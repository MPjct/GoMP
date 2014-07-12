package MySQLProtocol

type Packet_OK struct {
    Packet
    affected_rows uint64
    last_insert_id uint64
    status_flags uint16
    warnings uint16
    info string
    session_state_change string
}

func (packet Packet_OK) ToPacket(context Context) (data []byte) {
    var size uint64
    
    // Figure out size
    size = 4
    size += GetLengthEncodedIntegerSize(packet.affected_rows)
    size += GetLengthEncodedIntegerSize(packet.last_insert_id)
    if Has_Flag(context.client_capability, CLIENT_PROTOCOL_41) {
        size += 4
    } else if Has_Flag(context.client_capability, CLIENT_TRANSACTIONS) {
        size += 2
    }
    
    if Has_Flag(context.client_capability, CLIENT_SESSION_TRACK) {
        size += GetLengthEncodedStringSize(packet.info)
        if len(packet.session_state_change) > 0 {
            size += GetLengthEncodedStringSize(packet.session_state_change)
        }
    } else {
        size += uint64(len(packet.info))
    }
    
    data = make([]byte, 0, size)
    data = append(data, 0x00)
    data = append(data, BuildFixedLengthInteger3(uint32(size))...)
    data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    data = append(data, BuildLengthEncodedInteger(packet.affected_rows)...)
    data = append(data, BuildLengthEncodedInteger(packet.last_insert_id)...)
    
    if Has_Flag(context.client_capability, CLIENT_PROTOCOL_41) {
        data = append(data, BuildFixedLengthInteger2(packet.status_flags)...)
        data = append(data, BuildFixedLengthInteger2(packet.warnings)...)
    } else if Has_Flag(context.client_capability, CLIENT_TRANSACTIONS) {
        data = append(data, BuildFixedLengthInteger2(packet.status_flags)...)
    }
    
    if Has_Flag(context.client_capability, CLIENT_SESSION_TRACK) {
        data = append(data, BuildLengthEncodedString(packet.info)...)
        if len(packet.session_state_change) > 0 {
            data = append(data, BuildLengthEncodedString(packet.session_state_change)...)
        }
    } else {
        data = append(data, BuildFixedLengthString(packet.info)...)
    }
    
    return data
}
