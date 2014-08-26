package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_STMT_PREPARE_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
1c 00 00 00 16 53 45 4c    45 43 54 20 43 4f 4e 43    .....SELECT CONC
41 54 28 3f 2c 20 3f 29    20 41 53 20 63 6f 6c 31    AT(?, ?) AS col1
`)}, context: Context{}},
}

func Test_Packet_COM_STMT_PREPARE(t *testing.T) {
	var pkt Packet_COM_STMT_PREPARE

	for _, value := range COM_STMT_PREPARE_test_packets {
		pkt = Packet_COM_STMT_PREPARE{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_STMT_PREPARE_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_STMT_PREPARE
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_STMT_PREPARE{}
		COM_STMT_PREPARE_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_STMT_PREPARE_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_STMT_PREPARE_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_PREPARE{}
	pkt.FromPacket(context, COM_STMT_PREPARE_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_STMT_PREPARE_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_STMT_PREPARE{}
	pkt.FromPacket(context, COM_STMT_PREPARE_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
