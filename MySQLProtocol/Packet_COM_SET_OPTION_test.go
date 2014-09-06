package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_SET_OPTION_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
03 00 00 00 1b 01 00                                  .......
`)}, context: Context{}},
}

func Test_Packet_COM_SET_OPTION(t *testing.T) {
	var pkt Packet_COM_SET_OPTION

	for _, value := range COM_SET_OPTION_test_packets {
		pkt = Packet_COM_SET_OPTION{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_COM_SET_OPTION_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_COM_SET_OPTION
	for i := 0; i < b.N; i++ {
		pkt = Packet_COM_SET_OPTION{}
		COM_SET_OPTION_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, COM_SET_OPTION_test_packets[0].packet)
	}
}

func Benchmark_Packet_COM_SET_OPTION_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_SET_OPTION{}
	pkt.FromPacket(context, COM_SET_OPTION_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_COM_SET_OPTION_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_COM_SET_OPTION{}
	pkt.FromPacket(context, COM_SET_OPTION_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
