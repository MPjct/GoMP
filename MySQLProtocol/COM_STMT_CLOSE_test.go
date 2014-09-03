package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_STMT_CLOSE_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
05 00 00 00 19 01 00 00    00                         .........
`)}, context: Context{}},
}

func Test_Packet_COM_STMT_CLOSE(t *testing.T) {
	var pkt Packet_COM_STMT_CLOSE

	for _, value := range COM_STMT_CLOSE_test_packets {
		pkt = Packet_COM_STMT_CLOSE{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_STMT_CLOSE_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_STMT_CLOSE
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_STMT_CLOSE{}
		COM_STMT_CLOSE_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_STMT_CLOSE_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_STMT_CLOSE_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_CLOSE{}
	pkt.FromPacket(context, COM_STMT_CLOSE_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_STMT_CLOSE_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_CLOSE{}
	pkt.FromPacket(context, COM_STMT_CLOSE_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
