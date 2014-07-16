package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_HandshakeResponse41(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
54 00 00 01 8d a6 0f 00    00 00 00 01 08 00 00 00    T...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00 70 61 6d 00    14 ab 09 ee f6 bc b1 32    ....pam........2
3e 61 14 38 65 c0 99 1d    95 7d 75 d4 47 74 65 73    >a.8e....}u.Gtes
74 00 6d 79 73 71 6c 5f    6e 61 74 69 76 65 5f 70    t.mysql_native_p
61 73 73 77 6f 72 64 00                               assword.
`)}, context: Context{}},
		{packet: Proto{data: StringToPacket(`
b2 00 00 01 85 a2 1e 00    00 00 00 40 08 00 00 00    ...........@....
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00 72 6f 6f 74    00 14 22 50 79 a2 12 d4    ....root.."Py...
e8 82 e5 b3 f4 1a 97 75    6b c8 be db 9f 80 6d 79    .......uk.....my
73 71 6c 5f 6e 61 74 69    76 65 5f 70 61 73 73 77    sql_native_passw
6f 72 64 00 61 03 5f 6f    73 09 64 65 62 69 61 6e    ord.a._os.debian
36 2e 30 0c 5f 63 6c 69    65 6e 74 5f 6e 61 6d 65    6.0._client_name
08 6c 69 62 6d 79 73 71    6c 04 5f 70 69 64 05 32    .libmysql._pid.2
32 33 34 34 0f 5f 63 6c    69 65 6e 74 5f 76 65 72    2344._client_ver
73 69 6f 6e 08 35 2e 36    2e 36 2d 6d 39 09 5f 70    sion.5.6.6-m9._p
6c 61 74 66 6f 72 6d 06    78 38 36 5f 36 34 03 66    latform.x86_64.f
6f 6f 03 62 61 72                                     oo.bar
`)}, context: Context{}},
	}
	var pkt Packet_HandshakeResponse41

	for _, value := range values {
		pkt = Packet_HandshakeResponse41{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_HandshakeResponse41_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_HandshakeResponse41
	var packet = Proto{data: StringToPacket(`
54 00 00 01 8d a6 0f 00    00 00 00 01 08 00 00 00    T...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00 70 61 6d 00    14 ab 09 ee f6 bc b1 32    ....pam........2
3e 61 14 38 65 c0 99 1d    95 7d 75 d4 47 74 65 73    >a.8e....}u.Gtes
74 00 6d 79 73 71 6c 5f    6e 61 74 69 76 65 5f 70    t.mysql_native_p
61 73 73 77 6f 72 64 00                               assword.
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_HandshakeResponse41{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_HandshakeResponse41_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeResponse41{}
	var packet = Proto{data: StringToPacket(`
54 00 00 01 8d a6 0f 00    00 00 00 01 08 00 00 00    T...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00 70 61 6d 00    14 ab 09 ee f6 bc b1 32    ....pam........2
3e 61 14 38 65 c0 99 1d    95 7d 75 d4 47 74 65 73    >a.8e....}u.Gtes
74 00 6d 79 73 71 6c 5f    6e 61 74 69 76 65 5f 70    t.mysql_native_p
61 73 73 77 6f 72 64 00                               assword.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_HandshakeResponse41_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeResponse41{}
	var packet = Proto{data: StringToPacket(`
54 00 00 01 8d a6 0f 00    00 00 00 01 08 00 00 00    T...............
00 00 00 00 00 00 00 00    00 00 00 00 00 00 00 00    ................
00 00 00 00 70 61 6d 00    14 ab 09 ee f6 bc b1 32    ....pam........2
3e 61 14 38 65 c0 99 1d    95 7d 75 d4 47 74 65 73    >a.8e....}u.Gtes
74 00 6d 79 73 71 6c 5f    6e 61 74 69 76 65 5f 70    t.mysql_native_p
61 73 73 77 6f 72 64 00                               assword.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
