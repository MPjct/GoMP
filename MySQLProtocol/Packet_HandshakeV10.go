package MySQLProtocol

type Packet_HandshakeV10 struct {
    Packet
    
    protocol_version byte
    server_version string
    connection_id uint32
    auth_plugin_data_part_1 string
    capability uint32
    character_set byte
    status uint16
    auth_plugin_data_length uint8
    auth_plugin_data_part_2 string
    auth_plugin_name string
}

func (packet Packet_HandshakeV10) GetPacketSize(context Context) (size uint64) {
    size += 1
    return size
}

func (packet Packet_HandshakeV10) ToPacket(context Context) (data []byte) {
    size := packet.GetPacketSize(context)
    
    data = make([]byte, 0, size+4)
    return data
}

func (packet *Packet_HandshakeV10) FromPacket(context Context, data Proto) {
    data.GetFixedLengthInteger3()
    packet.sequence_id = data.GetFixedLengthInteger1()
    data.GetFixedLengthInteger1()
}
