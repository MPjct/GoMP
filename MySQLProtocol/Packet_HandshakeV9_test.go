package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_HandshakeV9(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
17 00 00 00 09 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00                   ..dvH@I-CJ.
`)}, context: Context{}},
	}
	var pkt Packet_HandshakeV9

	for _, value := range values {
		pkt = Packet_HandshakeV9{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_HandshakeV9_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_HandshakeV9
	var packet = Proto{data: StringToPacket(`
17 00 00 00 09 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00                   ..dvH@I-CJ.
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_HandshakeV9{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_HandshakeV9_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeV9{}
	var packet = Proto{data: StringToPacket(`
16 00 00 00 09 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00                   ..dvH@I-CJ.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_HandshakeV9_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeV9{}
	var packet = Proto{data: StringToPacket(`
17 00 00 00 09 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00                   ..dvH@I-CJ.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
