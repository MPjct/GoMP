package MySQLProtocol

type Packet_COM_STMT_EXECUTE struct {
	Packet

	statement_id          uint32
	flags                 uint8
	iteration_count       uint32
	nulls                 NULL_bitmap
	new_params_bound_flag uint8
	parameter_types       []uint16
	parameter_values      []string
}

func (packet Packet_COM_STMT_EXECUTE) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += 4
	size += 1
	size += 4

	if context.prepared_statements[packet.statement_id].num_params > 0 {
		size += packet.nulls.GetNullBitmapSize()
		size += 1

		if packet.new_params_bound_flag == 1 {

			for _, element := range packet.parameter_values {
				size += 2
				size += uint64(len(element))
			}
		}
	}
	return size
}

func (packet Packet_COM_STMT_EXECUTE) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)
	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	return data
}

func (packet *Packet_COM_STMT_EXECUTE) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()

}
