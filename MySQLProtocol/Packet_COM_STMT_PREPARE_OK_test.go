package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_STMT_PREPARE_OK_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
0c 00 00 01 00 01 00 00    00 01 00 02 00 00 00 00    ................

`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
0c 00 00 01 00 01 00 00    00 00 00 00 00 00 00 00
`)}, context: Context{}},
}

func Test_Packet_COM_STMT_PREPARE_OK(t *testing.T) {
	var pkt Packet_COM_STMT_PREPARE_OK

	for _, value := range COM_STMT_PREPARE_OK_test_packets {
		pkt = Packet_COM_STMT_PREPARE_OK{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_STMT_PREPARE_OK_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_STMT_PREPARE_OK
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_STMT_PREPARE_OK{}
		COM_STMT_PREPARE_OK_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_STMT_PREPARE_OK_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_STMT_PREPARE_OK_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_PREPARE_OK{}
	pkt.FromPacket(context, COM_STMT_PREPARE_OK_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_STMT_PREPARE_OK_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_PREPARE_OK{}
	pkt.FromPacket(context, COM_STMT_PREPARE_OK_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
