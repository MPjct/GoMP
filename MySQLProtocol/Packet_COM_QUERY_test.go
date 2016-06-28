package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_QUERY_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
21 00 00 00 03 73 65 6c    65 63 74 20 40 40 76 65    !....select @@ve
72 73 69 6f 6e 5f 63 6f    6d 6d 65 6e 74 20 6c 69    rsion_comment li
6d 69 74 20 31                                        mit 1
`)}, context: Context{}},
}

func Test_Packet_COM_QUERY(t *testing.T) {
	var pkt Packet_COM_QUERY

	for _, value := range COM_QUERY_test_packets {
		pkt = Packet_COM_QUERY{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_QUERY_FromPacket(b *testing.B) {
	context := COM_QUERY_test_packets[0].context
	packet := COM_QUERY_test_packets[0].packet
	pkt := Packet_COM_QUERY{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_COM_QUERY_GetPacketSize(b *testing.B) {
	context := COM_QUERY_test_packets[0].context
	pkt := Packet_COM_QUERY{}
	pkt.FromPacket(context, COM_QUERY_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_QUERY_ToPacket(b *testing.B) {
	context := COM_QUERY_test_packets[0].context
	pkt := Packet_COM_QUERY{}
	pkt.FromPacket(context, COM_QUERY_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
