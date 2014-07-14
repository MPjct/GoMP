package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_EOF(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
05 00 00 05 fe 00 00 02 00
`)}, context: Context{capability: CLIENT_PROTOCOL_41}},
	}
	var pkt Packet_EOF

	for _, value := range values {
		pkt = Packet_EOF{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_EOF_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_EOF
	var packet = Proto{data: StringToPacket(`
05 00 00 05 fe 00 00 02 00
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_EOF{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_EOF_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_EOF{}
	var packet = Proto{data: StringToPacket(`
05 00 00 05 fe 00 00 02 00
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_EOF_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_EOF{}
	var packet = Proto{data: StringToPacket(`
05 00 00 05 fe 00 00 02 00
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
