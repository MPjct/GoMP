package MySQLProtocol

type Packet_COM_CHANGE_USER struct {
	Packet

	character_set    uint16
	username         string
	auth_response    string
	database         string
	auth_plugin_name string
	attributes       Attributes
}

func (packet Packet_COM_CHANGE_USER) GetPacketSize(context Context) (size uint64) {
	size += 1
	size += GetNulTerminatedStringSize(packet.username)

	if Has_Flag(uint64(context.capability), CLIENT_SECURE_CONNECTION) {
		size += 1
		size += uint64(len(packet.auth_response))
	} else {
		size += GetNulTerminatedStringSize(packet.auth_response)
	}

	size += GetNulTerminatedStringSize(packet.database)

	if packet.character_set != 0x00 {
		size += 2

		if Has_Flag(uint64(context.capability), CLIENT_PLUGIN_AUTH) {
			size += GetNulTerminatedStringSize(packet.auth_plugin_name)
		}
		if Has_Flag(uint64(context.capability), CLIENT_CONNECT_ATTRS) {
			size += packet.attributes.GetAttributesSize()
		}
	}

	return size
}

func (packet Packet_COM_CHANGE_USER) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	data = append(data, BuildFixedLengthInteger1(COM_CHANGE_USER)...)
	data = append(data, BuildNulTerminatedString(packet.username)...)

	if Has_Flag(uint64(context.capability), CLIENT_SECURE_CONNECTION) {
		if len(packet.auth_response) > 255 {
			panic("packet.auth_response has to be less then 255 with this version of mysql. See https://dev.mysql.com/doc/internals/en/limitations.html")
		}
		data = append(data, BuildFixedLengthInteger1(uint8(len(packet.auth_response)))...)
		data = append(data, BuildFixedLengthString(packet.auth_response, uint(len(packet.auth_response)))...)
	} else {
		data = append(data, BuildNulTerminatedString(packet.auth_response)...)
	}

	data = append(data, BuildNulTerminatedString(packet.database)...)

	if packet.character_set != 0x00 {
		data = append(data, BuildFixedLengthInteger2(packet.character_set)...)

		if Has_Flag(uint64(context.capability), CLIENT_PLUGIN_AUTH) {
			data = append(data, BuildNulTerminatedString(packet.auth_plugin_name)...)
		}
		if Has_Flag(uint64(context.capability), CLIENT_CONNECT_ATTRS) {
			data = append(data, packet.attributes.BuildAttributes()...)
		}
	}

	return data
}

func (packet *Packet_COM_CHANGE_USER) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()

	data.GetFixedLengthInteger1()
	packet.username = data.GetNulTerminatedString()

	if Has_Flag(uint64(context.capability), CLIENT_SECURE_CONNECTION) {
		length := data.GetFixedLengthInteger1()
		packet.auth_response = data.GetFixedLengthString(uint(length))
	} else {
		packet.auth_response = data.GetNulTerminatedString()
	}

	packet.database = data.GetNulTerminatedString()

	if data.HasRemainingData() {
		packet.character_set = data.GetFixedLengthInteger2()

		if Has_Flag(uint64(context.capability), CLIENT_PLUGIN_AUTH) {
			packet.auth_plugin_name = data.GetNulTerminatedString()
		}
		if Has_Flag(uint64(context.capability), CLIENT_CONNECT_ATTRS) {
			packet.attributes = data.GetAttributes()
		}
	}
}
