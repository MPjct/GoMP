package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_CREATE_DB_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
05 00 00 00 05 74 65 73    74                         .....test
`)}, context: Context{}},
}

func Test_Packet_COM_CREATE_DB(t *testing.T) {
	var pkt Packet_COM_CREATE_DB
	for _, value := range COM_CREATE_DB_test_packets {
		pkt = Packet_COM_CREATE_DB{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_CREATE_DB_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_CREATE_DB
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_CREATE_DB{}
		COM_CREATE_DB_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_CREATE_DB_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_CREATE_DB_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_CREATE_DB{}
	pkt.FromPacket(context, COM_CREATE_DB_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_CREATE_DB_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_CREATE_DB{}
	pkt.FromPacket(context, COM_CREATE_DB_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
