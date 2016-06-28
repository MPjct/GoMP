package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_QUIT_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
01 00 00 00 01
`)}, context: Context{}},
}

func Test_Packet_COM_QUIT(t *testing.T) {
	var pkt Packet_COM_QUIT
	for _, value := range COM_QUIT_test_packets {
		pkt = Packet_COM_QUIT{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_QUIT_FromPacket(b *testing.B) {
	context := COM_QUIT_test_packets[0].context
	packet := COM_QUIT_test_packets[0].packet
	pkt := Packet_COM_QUIT{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_COM_QUIT_GetPacketSize(b *testing.B) {
	context := COM_QUIT_test_packets[0].context
	pkt := Packet_COM_QUIT{}
	pkt.FromPacket(context, COM_QUIT_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_QUIT_ToPacket(b *testing.B) {
	context := COM_QUIT_test_packets[0].context
	pkt := Packet_COM_QUIT{}
	pkt.FromPacket(context, COM_QUIT_test_packets[0].packet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
