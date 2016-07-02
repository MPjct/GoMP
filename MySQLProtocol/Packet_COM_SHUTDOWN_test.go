package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_SHUTDOWN_test_packets = []struct {
    packet  Proto
    context Context
}{
    {packet: Proto{data: StringToPacket(`
02 00 00 00 08 00
`)}, context: Context{}},
}

func Test_Packet_COM_SHUTDOWN(t *testing.T) {
    var pkt Packet_COM_SHUTDOWN
    for _, value := range COM_SHUTDOWN_test_packets {
        pkt = Packet_COM_SHUTDOWN{}
        pkt.FromPacket(value.context, value.packet)
        assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
    }

    short_pkt := Proto{data: StringToPacket("01 00 00 00 08")}
    pkt.FromPacket(COM_SHUTDOWN_test_packets[0].context, short_pkt)
    assert.Equal(t, pkt.ToPacket(COM_SHUTDOWN_test_packets[0].context), COM_SHUTDOWN_test_packets[0].packet.data, "")
}

func Benchmark_Packet_COM_SHUTDOWN_FromPacket(b *testing.B) {
    context := COM_SHUTDOWN_test_packets[0].context
    packet := COM_SHUTDOWN_test_packets[0].packet
    pkt := Packet_COM_SHUTDOWN{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        packet.offset = 0
        pkt.FromPacket(context, packet)
    }
}

func Benchmark_Packet_COM_SHUTDOWN_GetPacketSize(b *testing.B) {
    context := COM_SHUTDOWN_test_packets[0].context
    pkt := Packet_COM_SHUTDOWN{}
    pkt.FromPacket(context, COM_SHUTDOWN_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.GetPacketSize(context)
    }
}

func Benchmark_Packet_COM_SHUTDOWN_ToPacket(b *testing.B) {
    context := COM_SHUTDOWN_test_packets[0].context
    pkt := Packet_COM_SHUTDOWN{}
    pkt.FromPacket(context, COM_SHUTDOWN_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.ToPacket(context)
    }
}
