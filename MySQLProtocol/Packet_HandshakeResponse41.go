package MySQLProtocol

type Packet_HandshakeResponse41 struct {
	Packet

	capability       uint32
	max_packet_size  uint32
	character_set    byte
	username         string
	auth_response    string
	database         string
	auth_plugin_name string
	attributes       Attributes
}

func (packet Packet_HandshakeResponse41) GetPacketSize(context Context) (size uint64) {
	size += 4
	size += 4
	size += 1
	size += 23
	size += GetNulTerminatedStringSize(packet.username)
	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA) {
		size += GetLengthEncodedStringSize(packet.auth_response)
	} else if Has_Flag(uint64(packet.capability), CLIENT_SECURE_CONNECTION) {
		size += 1
		size += uint64(len(packet.auth_response))
	} else {
		size += GetNulTerminatedStringSize(packet.auth_response)
	}

	if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_WITH_DB) {
		size += GetNulTerminatedStringSize(packet.database)
	}

	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH) {
		size += GetNulTerminatedStringSize(packet.auth_plugin_name)
	}

	if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_ATTRS) {
		size += packet.attributes.GetAttributesSize()
	}
	return size
}

func (packet Packet_HandshakeResponse41) ToPacket(context Context) (data []byte) {
	size := packet.GetPacketSize(context)

	data = make([]byte, 0, size+4)

	data = append(data, BuildFixedLengthInteger3(uint32(size))...)
	data = append(data, BuildFixedLengthInteger1(packet.sequence_id)...)

	data = append(data, BuildFixedLengthInteger4(packet.capability)...)
	data = append(data, BuildFixedLengthInteger4(packet.max_packet_size)...)
	data = append(data, BuildFixedLengthInteger1(packet.character_set)...)
	data = append(data, BuildFixedLengthString("", 23)...)
	data = append(data, BuildNulTerminatedString(packet.username)...)

	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA) {
		data = append(data, BuildLengthEncodedString(packet.auth_response)...)
	} else if Has_Flag(uint64(packet.capability), CLIENT_SECURE_CONNECTION) {
		if len(packet.auth_response) > 255 {
			panic("packet.auth_response has to be less then 255 with this version of mysql. See https://dev.mysql.com/doc/internals/en/limitations.html")
		}
		data = append(data, BuildFixedLengthInteger1(uint8(len(packet.auth_response)))...)
		data = append(data, BuildFixedLengthString(packet.auth_response, uint(len(packet.auth_response)))...)
	} else {
		data = append(data, BuildNulTerminatedString(packet.auth_response)...)
	}

	if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_WITH_DB) {
		data = append(data, BuildNulTerminatedString(packet.database)...)
	}

	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH) {
		data = append(data, BuildNulTerminatedString(packet.auth_plugin_name)...)
	}

	if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_ATTRS) {
		data = append(data, packet.attributes.BuildAttributes()...)
	}

	return data
}

func (packet *Packet_HandshakeResponse41) FromPacket(context Context, data Proto) {
	data.GetFixedLengthInteger3()
	packet.sequence_id = data.GetFixedLengthInteger1()

	packet.capability = data.GetFixedLengthInteger4()
	packet.max_packet_size = data.GetFixedLengthInteger4()
	packet.character_set = data.GetFixedLengthInteger1()
	data.GetFixedLengthString(23)
	packet.username = data.GetNulTerminatedString()
	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA) {
		packet.auth_response = data.GetLengthEncodedString()
	} else if Has_Flag(uint64(packet.capability), CLIENT_SECURE_CONNECTION) {
		length := data.GetFixedLengthInteger1()
		packet.auth_response = data.GetFixedLengthString(uint(length))
	} else {
		packet.auth_response = data.GetNulTerminatedString()
	}

	if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_WITH_DB) {
		packet.database = data.GetNulTerminatedString()
	}

	if Has_Flag(uint64(packet.capability), CLIENT_PLUGIN_AUTH) {
		packet.auth_plugin_name = data.GetNulTerminatedString()
	}

	if Has_Flag(uint64(packet.capability), CLIENT_CONNECT_ATTRS) {
		packet.attributes = data.GetAttributes()
	}
}
