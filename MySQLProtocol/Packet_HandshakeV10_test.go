package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_HandshakeV10(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
36 00 00 00 0a 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00 ff f7 08 02 00    ..dvH@I-CJ......
00 00 00 00 00 00 00 00    00 00 00 00 00 2a 34 64    .............*4d
7c 63 5a 77 6b 34 5e 5d    3a 00                      |cZwk4^]:.
`)}, context: Context{}},

		{packet: Proto{data: StringToPacket(`
50 00 00 00 0a 35 2e 36    2e 34 2d 6d 37 2d 6c 6f    P....5.6.4-m7-lo
67 00 56 0a 00 00 52 42    33 76 7a 26 47 72 00 ff    g.V...RB3vz&Gr..
ff 08 02 00 0f c0 15 00    00 00 00 00 00 00 00 00    ................
00 2b 79 44 26 2f 5a 5a    33 30 35 5a 47 00 6d 79    .+yD&/ZZ305ZG.my
73 71 6c 5f 6e 61 74 69    76 65 5f 70 61 73 73 77    sql_native_passw
6f 72 64 00                                           ord
`)}, context: Context{}},

	}
	var pkt Packet_HandshakeV10

	for _, value := range values {
		pkt = Packet_HandshakeV10{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_HandshakeV10_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_HandshakeV10
	var packet = Proto{data: StringToPacket(`
36 00 00 00 0a 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00 ff f7 08 02 00    ..dvH@I-CJ......
00 00 00 00 00 00 00 00    00 00 00 00 00 2a 34 64    .............*4d
7c 63 5a 77 6b 34 5e 5d    3a 00                      |cZwk4^]:.
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_HandshakeV10{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_HandshakeV10_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeV10{}
	var packet = Proto{data: StringToPacket(`
36 00 00 00 0a 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00 ff f7 08 02 00    ..dvH@I-CJ......
00 00 00 00 00 00 00 00    00 00 00 00 00 2a 34 64    .............*4d
7c 63 5a 77 6b 34 5e 5d    3a 00                      |cZwk4^]:.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_HandshakeV10_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeV10{}
	var packet = Proto{data: StringToPacket(`
36 00 00 00 0a 35 2e 35    2e 32 2d 6d 32 00 0b 00    6....5.5.2-m2...
00 00 64 76 48 40 49 2d    43 4a 00 ff f7 08 02 00    ..dvH@I-CJ......
00 00 00 00 00 00 00 00    00 00 00 00 00 2a 34 64    .............*4d
7c 63 5a 77 6b 34 5e 5d    3a 00                      |cZwk4^]:.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
