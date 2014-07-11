package MySQLProtocol

type Packet struct {
	data   []byte
	offset uint
}

func (packet Packet) HasRemainingData() bool {
	return uint(len(packet.data))-packet.offset > 0
}
