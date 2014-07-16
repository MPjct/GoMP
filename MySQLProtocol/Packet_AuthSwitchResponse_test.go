package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_AuthSwitchResponse(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
09 00 00 03 5c 49 4d 5e    4e 58 4f 47 00             ....\IM^NXOG.
`)}, context: Context{}},
		{packet: Proto{data: StringToPacket(`
14 00 00 03 f4 17 96 1f    79 f3 ac 10 0b da a6 b3    ........y.......
b5 c2 0e ab 59 85 ff b8                               ....Y...
`)}, context: Context{}},
	}
	var pkt Packet_AuthSwitchResponse

	for _, value := range values {
		pkt = Packet_AuthSwitchResponse{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_AuthSwitchResponse_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_AuthSwitchResponse
	var packet = Proto{data: StringToPacket(`
09 00 00 03 5c 49 4d 5e    4e 58 4f 47 00             ....\IM^NXOG.
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_AuthSwitchResponse{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_AuthSwitchResponse_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_AuthSwitchResponse{}
	var packet = Proto{data: StringToPacket(`
09 00 00 03 5c 49 4d 5e    4e 58 4f 47 00             ....\IM^NXOG.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_AuthSwitchResponse_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_AuthSwitchResponse{}
	var packet = Proto{data: StringToPacket(`
09 00 00 03 5c 49 4d 5e    4e 58 4f 47 00             ....\IM^NXOG.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
