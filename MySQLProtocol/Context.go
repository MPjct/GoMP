package MySQLProtocol

type Context struct {
    capability uint64
    prepared_statements map[uint32]Packet_COM_STMT_PREPARE_OK
}
