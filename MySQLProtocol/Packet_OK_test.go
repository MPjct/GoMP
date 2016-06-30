package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var Packet_OK_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
07 00 00 02 00 00 00 02    00 00 00                   ...........
`)}, context: Context{capability: CLIENT_PROTOCOL_41}},

	{packet: Proto{data: StringToPacket(`
07 00 00 02 00 00 00 02    00 00 00                   ...........
`)}, context: Context{capability: CLIENT_TRANSACTIONS}},

	{packet: Proto{data: StringToPacket(`
0d 00 00 02 00 00 00 02    40 00 00 00 04 03 02 01    ........@.......
31                                                    1
`)}, context: Context{capability: CLIENT_PROTOCOL_41 | CLIENT_TRANSACTIONS | CLIENT_SESSION_TRACK}},

	{packet: Proto{data: StringToPacket(`
10 00 00 02 00 00 00 02    40 00 00 00 07 01 05 04    ........@.......
74 65 73 74                                           test
`)}, context: Context{capability: CLIENT_PROTOCOL_41 | CLIENT_TRANSACTIONS | CLIENT_SESSION_TRACK}},

	{packet: Proto{data: StringToPacket(`
1d 00 00 01 00 00 00 00    40 00 00 00 14 00 0f 0a    ........@.......
61 75 74 6f 63 6f 6d 6d    69 74 03 4f 46 46 02 01    autocommit.OFF..
31                                                    1
`)}, context: Context{capability: CLIENT_PROTOCOL_41 | CLIENT_TRANSACTIONS | CLIENT_SESSION_TRACK}},

	{packet: Proto{data: StringToPacket(`
13 00 00 01 00 00 00 00    40 00 00 00 0a 01 05 04    ........@.......
74 65 73 74 02 01 31                                  test..1
`)}, context: Context{capability: CLIENT_PROTOCOL_41 | CLIENT_TRANSACTIONS | CLIENT_SESSION_TRACK}},
}

func Test_Packet_OK(t *testing.T) {
	var pkt Packet_OK
	for _, value := range Packet_OK_test_packets {
		pkt = Packet_OK{}
		pkt.FromPacket(value.context, value.packet)
		if !assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "") {
			DumpPacket(value.packet.data)
			DumpPacket(pkt.ToPacket(value.context))
		}
	}
}

func Benchmark_Packet_OK_FromPacket(b *testing.B) {
	context := Packet_OK_test_packets[0].context
	packet := Packet_OK_test_packets[0].packet
	pkt := Packet_OK{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_OK_GetPacketSize(b *testing.B) {
	context := Packet_OK_test_packets[0].context
	packet := Packet_OK_test_packets[0].packet
	pkt := Packet_OK{}
	pkt.FromPacket(context, packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_OK_ToPacket(b *testing.B) {
	context := Packet_OK_test_packets[0].context
	packet := Packet_OK_test_packets[0].packet
	pkt := Packet_OK{}
	pkt.FromPacket(context, packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
