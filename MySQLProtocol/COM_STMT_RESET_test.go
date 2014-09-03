package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_STMT_RESET_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
05 00 00 00 1a 01 00 00    00                         .........
`)}, context: Context{}},
}

func Test_Packet_COM_STMT_RESET(t *testing.T) {
	var pkt Packet_COM_STMT_RESET

	for _, value := range COM_STMT_RESET_test_packets {
		pkt = Packet_COM_STMT_RESET{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_STMT_RESET_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_STMT_RESET
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_STMT_RESET{}
		COM_STMT_RESET_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_STMT_RESET_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_STMT_RESET_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_RESET{}
	pkt.FromPacket(context, COM_STMT_RESET_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_STMT_RESET_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_RESET{}
	pkt.FromPacket(context, COM_STMT_RESET_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
