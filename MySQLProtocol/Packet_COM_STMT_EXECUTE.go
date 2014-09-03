package MySQLProtocol

type Packet_COM_STMT_EXECUTE struct {
	Packet

	statement_id          uint32
	flags                 uint8
	iteration_count       uint32
	nulls                 NULL_bitmap
	new_params_bound_flag uint8
    parameters            []Field_interface
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

			for _, parameter := range packet.parameters {
				size += 2
				size += parameter.Size()
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
    
    data = append(data, BuildFixedLengthInteger4(packet.statement_id)...)
    data = append(data, BuildFixedLengthInteger1(packet.flags)...)
    data = append(data, BuildFixedLengthInteger4(packet.iteration_count)...)
    data = append(data, packet.nulls.BuildNullBitmap()...)
    data = append(data, BuildFixedLengthInteger1(packet.new_params_bound_flag)...)

    for _, parameter := range packet.parameters {
        data = append(data, BuildFixedLengthInteger2(uint16(parameter.GetType()))...)
    }
    
    for _, parameter := range packet.parameters {
        data = append(data, parameter.Build()...)
    }

	return data
}

func (packet *Packet_COM_STMT_EXECUTE) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()
    
    num_params := context.prepared_statements[packet.statement_id].num_params
    
    packet.statement_id = data.GetFixedLengthInteger4()
    packet.flags = data.GetFixedLengthInteger1()
    packet.iteration_count = data.GetFixedLengthInteger4()
    packet.nulls = data.GetNullBitmap(num_params, 0)
    packet.new_params_bound_flag = data.GetFixedLengthInteger1()
    
    packet.parameters = make([]Field_interface, 0, num_params)
    parameters := make([]byte, 0, num_params)
    
    for i := uint16(0); i < num_params; i++ {
        parameters = append(parameters, byte(data.GetFixedLengthInteger2()))
    }
    
    for _, parameter := range parameters {
        packet.parameters = append(packet.parameters, data.GetField(parameter))
    }
}
