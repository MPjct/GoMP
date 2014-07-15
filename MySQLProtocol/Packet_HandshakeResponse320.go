package MySQLProtocol

type Packet_HandshakeResponse320 struct {
	Packet
    
    capability uint16
    max_packet_size uint32
    username string
    auth_response string
    database string
}

func (packet Packet_HandshakeResponse320) GetPacketSize(context Context) (size uint64) {
	size += 2
    size += 3
	size += GetNulTerminatedStringSize(packet.username)
    
    if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_WITH_DB) {
        size += GetNulTerminatedStringSize(packet.auth_response)
        size += GetNulTerminatedStringSize(packet.database)
    } else {
        size += GetFixedLengthStringSize(packet.auth_response)
    }
    
	return size
}

func (packet Packet_HandshakeResponse320) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
    data = append(data, BuildFixedLengthInteger2(packet.capability)...)
    data = append(data, BuildFixedLengthInteger3(packet.max_packet_size)...)
    data = append(data, BuildNulTerminatedString(packet.username)...)
    if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_WITH_DB) {
        data = append(data, BuildNulTerminatedString(packet.auth_response)...)
        data = append(data, BuildNulTerminatedString(packet.database)...)
    } else {
        data = append(data, BuildFixedLengthString(packet.auth_response)...)
    }
    
	return data
}

func (packet *Packet_HandshakeResponse320) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
    
    packet.capability = data.GetFixedLengthInteger2()
    packet.max_packet_size = data.GetFixedLengthInteger3()
    packet.username = data.GetNulTerminatedString()
    
    if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_WITH_DB) {
        packet.auth_response = data.GetNulTerminatedString()
        packet.database = data.GetNulTerminatedString()
    } else {
        packet.auth_response = data.GetFixedLengthString()
    }
}
