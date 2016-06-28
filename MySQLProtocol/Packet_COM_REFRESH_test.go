package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_REFRESH_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 01
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 02
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 04
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 08
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 10
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 20
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 40
`)}, context: Context{}},
	{packet: Proto{data: StringToPacket(`
02 00 00 00 07 80
`)}, context: Context{}},
}

func Test_Packet_COM_REFRESH(t *testing.T) {
	var pkt Packet_COM_REFRESH
	for _, value := range COM_REFRESH_test_packets {
		pkt = Packet_COM_REFRESH{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_REFRESH_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_REFRESH
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_REFRESH{}
		COM_REFRESH_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_REFRESH_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_REFRESH_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_REFRESH{}
	pkt.FromPacket(context, COM_REFRESH_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_REFRESH_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_REFRESH{}
	pkt.FromPacket(context, COM_REFRESH_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
