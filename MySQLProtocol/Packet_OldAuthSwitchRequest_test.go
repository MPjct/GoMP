package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_OldAuthSwitchRequest(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
01 00 00 02 fe
`)}, context: Context{}},
	}
	var pkt Packet_OldAuthSwitchRequest

	for _, value := range values {
		pkt = Packet_OldAuthSwitchRequest{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_OldAuthSwitchRequest_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_OldAuthSwitchRequest
	var packet = Proto{data: StringToPacket(`
01 00 00 02 fe
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_OldAuthSwitchRequest{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_OldAuthSwitchRequest_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_OldAuthSwitchRequest{}
	var packet = Proto{data: StringToPacket(`
01 00 00 02 fe
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_OldAuthSwitchRequest_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_OldAuthSwitchRequest{}
	var packet = Proto{data: StringToPacket(`
01 00 00 02 fe
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
