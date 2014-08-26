package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var ColumnDefinition41_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
17 00 00 02 03 64 65 66    00 00 00 01 3f 00 0c 3f    .....def....?..?
00 00 00 00 00 fd 80 00    00 00 00

`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
1a 00 00 05 03 64 65 66    00 00 00 04 63 6f 6c 31
00 0c 3f 00 00 00 00 00    fd 80 00 1f 00 00
`)}, context: Context{}},
}

func Test_Packet_ColumnDefinition41(t *testing.T) {
	var pkt Packet_ColumnDefinition41

	for _, value := range ColumnDefinition41_test_packets {
		pkt = Packet_ColumnDefinition41{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_ColumnDefinition41_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_ColumnDefinition41
	for i := 0; i < b.N; i++ {
		pkt = Packet_ColumnDefinition41{}
		ColumnDefinition41_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, ColumnDefinition41_test_packets[0].packet)
	}
}

func Benchmark_Packet_ColumnDefinition41_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_ColumnDefinition41{}
	pkt.FromPacket(context, ColumnDefinition41_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_ColumnDefinition41_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_ColumnDefinition41{}
	pkt.FromPacket(context, ColumnDefinition41_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
