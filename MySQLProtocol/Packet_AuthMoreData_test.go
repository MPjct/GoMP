package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var AuthMoreData_test_packets = []struct {
	packet  Proto
	context Context
}{

	{packet: Proto{data: StringToPacket(`
05 00 00 00 01 74 65 73 74
`)}, context: Context{}},
}

func Test_Packet_AuthMoreData(t *testing.T) {
	var pkt Packet_AuthMoreData
	for _, value := range AuthMoreData_test_packets {
		pkt = Packet_AuthMoreData{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_AuthMoreData_FromPacket(b *testing.B) {
	context := AuthMoreData_test_packets[0].context
	packet := AuthMoreData_test_packets[0].packet
	pkt := Packet_AuthMoreData{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_AuthMoreData_GetPacketSize(b *testing.B) {
	context := AuthMoreData_test_packets[0].context
	pkt := Packet_AuthMoreData{}
	pkt.FromPacket(context, AuthMoreData_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_AuthMoreData_ToPacket(b *testing.B) {
	context := AuthMoreData_test_packets[0].context
	pkt := Packet_AuthMoreData{}
	pkt.FromPacket(context, AuthMoreData_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
