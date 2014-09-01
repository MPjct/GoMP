package MySQLProtocol

type Packet_COM_SHUTDOWN struct {
	Packet
	shutdown_type uint8
}

func (packet Packet_COM_SHUTDOWN) GetPacketSize(context Context) (size uint64) {
	size += 2
	return size
}

func (packet Packet_COM_SHUTDOWN) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)
	data = append(data, COM_SHUTDOWN)
	data = append(data, packet.shutdown_type)
	return data
}

func (packet *Packet_COM_SHUTDOWN) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
	data.GetFixedLengthInteger1()
	if data.HasRemainingData() {
		packet.shutdown_type = data.GetFixedLengthInteger1()
	} else {
		packet.shutdown_type = 0x00
	}
}
