package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var COM_PROCESS_KILL_test_packets = []struct {
    packet  Proto
    context Context
}{
    {packet: Proto{data: StringToPacket(`
05 00 00 00 0c 12 34 56 78
`)}, context: Context{}},
}

func Test_Packet_COM_PROCESS_KILL(t *testing.T) {
    var pkt Packet_COM_PROCESS_KILL
    for _, value := range COM_PROCESS_KILL_test_packets {
        pkt = Packet_COM_PROCESS_KILL{}
        pkt.FromPacket(value.context, value.packet)
        assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
    }
}

func Benchmark_Packet_COM_PROCESS_KILL_FromPacket(b *testing.B) {
    context := COM_PROCESS_KILL_test_packets[0].context
    packet := COM_PROCESS_KILL_test_packets[0].packet
    pkt := Packet_COM_PROCESS_KILL{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        packet.offset = 0
        pkt.FromPacket(context, packet)
    }
}

func Benchmark_Packet_COM_PROCESS_KILL_GetPacketSize(b *testing.B) {
    context := COM_PROCESS_KILL_test_packets[0].context
    pkt := Packet_COM_PROCESS_KILL{}
    pkt.FromPacket(context, COM_PROCESS_KILL_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.GetPacketSize(context)
    }
}

func Benchmark_Packet_COM_PROCESS_KILL_ToPacket(b *testing.B) {
    context := COM_PROCESS_KILL_test_packets[0].context
    pkt := Packet_COM_PROCESS_KILL{}
    pkt.FromPacket(context, COM_PROCESS_KILL_test_packets[0].packet)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pkt.ToPacket(context)
    }
}
