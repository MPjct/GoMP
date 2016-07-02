package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_PING_test_packets = []struct {
    packet  Proto
    context Context
}{
    {packet: Proto{data: StringToPacket(`
01 00 00 00 0e
`)}, context: Context{}},
}

func Test_Packet_COM_PING(t *testing.T) {
    var pkt Packet_COM_PING
    for _, value := range COM_PING_test_packets {
        pkt = Packet_COM_PING{}
        pkt.FromPacket(value.context, value.packet)
        assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
    }
}

func Benchmark_Packet_COM_PING_FromPacket(b *testing.B) {
    context := COM_PING_test_packets[0].context
    packet := COM_PING_test_packets[0].packet
    pkt := Packet_COM_PING{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        packet.offset = 0
        pkt.FromPacket(context, packet)
    }
}

func Benchmark_Packet_COM_PING_GetPacketSize(b *testing.B) {
    context := COM_PING_test_packets[0].context
    pkt := Packet_COM_PING{}
    pkt.FromPacket(context, COM_PING_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.GetPacketSize(context)
    }
}

func Benchmark_Packet_COM_PING_ToPacket(b *testing.B) {
    context := COM_PING_test_packets[0].context
    pkt := Packet_COM_PING{}
    pkt.FromPacket(context, COM_PING_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.ToPacket(context)
    }
}
