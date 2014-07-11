package MySQLProtocol

type Proto struct {
	data   []byte
	offset uint
}

func (packet Proto) HasRemainingData() bool {
	return uint(len(packet.data))-packet.offset > 0
}
