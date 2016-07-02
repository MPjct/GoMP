package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_FIELD_LIST_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
0a 00 00 00 04 74 65 73 74 00 74 65 73 74
`)}, context: Context{}},
}

func Test_Packet_COM_FIELD_LIST(t *testing.T) {
	var pkt Packet_COM_FIELD_LIST
	for _, value := range COM_FIELD_LIST_test_packets {
		pkt = Packet_COM_FIELD_LIST{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_FIELD_LIST_FromPacket(b *testing.B) {
	context := COM_FIELD_LIST_test_packets[0].context
	packet := COM_FIELD_LIST_test_packets[0].packet
	pkt := Packet_COM_FIELD_LIST{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_COM_FIELD_LIST_GetPacketSize(b *testing.B) {
	context := COM_FIELD_LIST_test_packets[0].context
	pkt := Packet_COM_FIELD_LIST{}
	pkt.FromPacket(context, COM_FIELD_LIST_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_FIELD_LIST_ToPacket(b *testing.B) {
	context := COM_FIELD_LIST_test_packets[0].context
	pkt := Packet_COM_FIELD_LIST{}
	pkt.FromPacket(context, COM_FIELD_LIST_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
