package Binlog

import "github.com/MPjct/GoMP/MySQLProtocol"

// https://dev.mysql.com/doc/internals/en/binlog-event-header.html

type EventHeader struct {
	MySQLProtocol.Packet
	timestamp      uint32
	event_type     uint8
	server_id      uint32
	size           uint32
	next_event_pos uint32
	flags          uint16
}

func (event EventHeader) GetPacketSize(context Context) (size uint64) {
	if context.binlog_version == Binlog_Version_1 {
		return 13
	}
	return 19
}

func (event EventHeader) ToPacket(context Context) (data []byte) {
	size := event.GetPacketSize(context)

	data = make([]byte, 0, size)
	data = append(data, MySQLProtocol.BuildFixedLengthInteger4(event.timestamp)...)
	data = append(data, MySQLProtocol.BuildFixedLengthInteger1(event.event_type)...)
	data = append(data, MySQLProtocol.BuildFixedLengthInteger4(event.server_id)...)
	data = append(data, MySQLProtocol.BuildFixedLengthInteger4(event.size)...)

	if context.binlog_version != Binlog_Version_1 {
		data = append(data, MySQLProtocol.BuildFixedLengthInteger4(event.next_event_pos)...)
		data = append(data, MySQLProtocol.BuildFixedLengthInteger2(event.flags)...)
	}

	return data
}

func (event *EventHeader) FromPacket(context Context, data MySQLProtocol.Proto) {

	event.timestamp = data.GetFixedLengthInteger4()
	event.event_type = data.GetFixedLengthInteger1()
	event.server_id = data.GetFixedLengthInteger4()
	event.size = data.GetFixedLengthInteger4()

	if context.binlog_version != Binlog_Version_1 {
		event.next_event_pos = data.GetFixedLengthInteger4()
		event.flags = data.GetFixedLengthInteger2()
	}
}
