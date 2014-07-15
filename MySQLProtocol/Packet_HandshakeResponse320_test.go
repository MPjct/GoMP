package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_HandshakeResponse320(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
11 00 00 01 85 24 00 00    00 6f 6c 64 00 47 44 53    .....$...old.GDS
43 51 59 52 5f                                        CQYR_
`)}, context: Context{}},
	}
	var pkt Packet_HandshakeResponse320

	for _, value := range values {
		pkt = Packet_HandshakeResponse320{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_HandshakeResponse320_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_HandshakeResponse320
	var packet = Proto{data: StringToPacket(`
11 00 00 01 85 24 00 00    00 6f 6c 64 00 47 44 53    .....$...old.GDS
43 51 59 52 5f                                        CQYR_
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_HandshakeResponse320{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_HandshakeResponse320_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeResponse320{}
	var packet = Proto{data: StringToPacket(`
11 00 00 01 85 24 00 00    00 6f 6c 64 00 47 44 53    .....$...old.GDS
43 51 59 52 5f                                        CQYR_
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_HandshakeResponse320_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_HandshakeResponse320{}
	var packet = Proto{data: StringToPacket(`
11 00 00 01 85 24 00 00    00 6f 6c 64 00 47 44 53    .....$...old.GDS
43 51 59 52 5f                                        CQYR_
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
