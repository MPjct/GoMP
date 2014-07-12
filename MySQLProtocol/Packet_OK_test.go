package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_OK(t *testing.T) {
	var values = []struct {
		packet  Proto
        context Context
	}{
		{packet: Proto{data: []byte{0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00}},
        context: Context{client_capability: CLIENT_PROTOCOL_41}},
	}
    var pkt Packet_OK

	for _, value := range values {
        pkt = Packet_OK{}
        pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_OK_From_Packet(b *testing.B) {
    context := Context{}
    var pkt Packet_OK
    var packet = Proto{data: []byte{0x00}}
	for i := 0; i < b.N; i++ {
        pkt = Packet_OK{}
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_OK_To_Packet(b *testing.B) {
    pkt := Packet_OK{}
    context := Context{}
    var packet = Proto{data: []byte{0x00}}
    pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
