package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_CONNECT_test_packets = []struct {
    packet  Proto
    context Context
}{
    {packet: Proto{data: StringToPacket(`
01 00 00 00 0b
`)}, context: Context{}},
}

func Test_Packet_COM_CONNECT(t *testing.T) {
    var pkt Packet_COM_CONNECT
    for _, value := range COM_CONNECT_test_packets {
        pkt = Packet_COM_CONNECT{}
        pkt.FromPacket(value.context, value.packet)
        assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
    }
}

func Benchmark_Packet_COM_CONNECT_FromPacket(b *testing.B) {
    context := COM_CONNECT_test_packets[0].context
    packet := COM_CONNECT_test_packets[0].packet
    pkt := Packet_COM_CONNECT{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        packet.offset = 0
        pkt.FromPacket(context, packet)
    }
}

func Benchmark_Packet_COM_CONNECT_GetPacketSize(b *testing.B) {
    context := COM_CONNECT_test_packets[0].context
    pkt := Packet_COM_CONNECT{}
    pkt.FromPacket(context, COM_CONNECT_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.GetPacketSize(context)
    }
}

func Benchmark_Packet_COM_CONNECT_ToPacket(b *testing.B) {
    context := COM_CONNECT_test_packets[0].context
    pkt := Packet_COM_CONNECT{}
    pkt.FromPacket(context, COM_CONNECT_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.ToPacket(context)
    }
}
