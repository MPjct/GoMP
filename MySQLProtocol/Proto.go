package MySQLProtocol

type Proto struct {
	data   []byte
	offset uint
}

func (proto *Proto) HasRemainingData() bool {
	return uint(len(proto.data))-proto.offset > 0
}
