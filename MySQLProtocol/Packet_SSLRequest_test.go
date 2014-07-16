package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_SSLRequest(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
20 00 00 01 05 ae 03 00    00 00 00 01 08 00 00 00     ...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00                                           ....
`)}, context: Context{}},
	}
	var pkt Packet_SSLRequest

	for _, value := range values {
		pkt = Packet_SSLRequest{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_SSLRequest_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_SSLRequest
	var packet = Proto{data: StringToPacket(`
20 00 00 01 05 ae 03 00    00 00 00 01 08 00 00 00     ...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00                                           ....
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_SSLRequest{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_SSLRequest_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_SSLRequest{}
	var packet = Proto{data: StringToPacket(`
20 00 00 01 05 ae 03 00    00 00 00 01 08 00 00 00     ...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00                                           ....
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_SSLRequest_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_SSLRequest{}
	var packet = Proto{data: StringToPacket(`
20 00 00 01 05 ae 03 00    00 00 00 01 08 00 00 00     ...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00                                           ....
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
