package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_ERR(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: []byte{
			0x17, 0x00, 0x00, 0x01, 0xff, 0x48, 0x04, 0x23,
			0x48, 0x59, 0x30, 0x30, 0x30, 0x4e, 0x6f, 0x20,
			0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x75,
			0x73, 0x65, 0x64,
		}},
			context: Context{capability: CLIENT_PROTOCOL_41}},
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
	var packet = Proto{data: []byte{
		0x17, 0x00, 0x00, 0x01, 0xff, 0x48, 0x04, 0x23,
		0x48, 0x59, 0x30, 0x30, 0x30, 0x4e, 0x6f, 0x20,
		0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x75,
		0x73, 0x65, 0x64,
	}}
	for i := 0; i < b.N; i++ {
		pkt = Packet_ERR{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_ERR_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_ERR{}
	var packet = Proto{data: []byte{
		0x17, 0x00, 0x00, 0x01, 0xff, 0x48, 0x04, 0x23,
		0x48, 0x59, 0x30, 0x30, 0x30, 0x4e, 0x6f, 0x20,
		0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x75,
		0x73, 0x65, 0x64,
	}}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_ERR_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_ERR{}
	var packet = Proto{data: []byte{
		0x17, 0x00, 0x00, 0x01, 0xff, 0x48, 0x04, 0x23,
		0x48, 0x59, 0x30, 0x30, 0x30, 0x4e, 0x6f, 0x20,
		0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x20, 0x75,
		0x73, 0x65, 0x64,
	}}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
