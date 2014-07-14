package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_ERR(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
17 00 00 01 ff 48 04 23    48 59 30 30 30 4e 6f 20    .....H.#HY000No
74 61 62 6c 65 73 20 75    73 65 64                   tables used
`)}, context: Context{capability: CLIENT_PROTOCOL_41}},
	}
	var pkt Packet_ERR

	for _, value := range values {
		pkt = Packet_ERR{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_ERR_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_ERR
	var packet = Proto{data: StringToPacket(`
17 00 00 01 ff 48 04 23    48 59 30 30 30 4e 6f 20    .....H.#HY000No
74 61 62 6c 65 73 20 75    73 65 64                   tables used
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_ERR{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_ERR_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_ERR{}
	var packet = Proto{data: StringToPacket(`
17 00 00 01 ff 48 04 23    48 59 30 30 30 4e 6f 20    .....H.#HY000No
74 61 62 6c 65 73 20 75    73 65 64                   tables used
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_ERR_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_ERR{}
	var packet = Proto{data: StringToPacket(`
17 00 00 01 ff 48 04 23    48 59 30 30 30 4e 6f 20    .....H.#HY000No
74 61 62 6c 65 73 20 75    73 65 64                   tables used
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
