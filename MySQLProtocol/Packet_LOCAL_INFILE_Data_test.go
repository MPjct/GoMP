package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var LOCAL_INFILE_Data_test_packets = []struct {
	packet  Proto
	context Context
}{
	{packet: Proto{data: StringToPacket(`
0b 00 00 01 2f 65 74 63   2f 70 61 73 73 77 64    ..../etc/passwd
`)}, context: Context{}},
}

func Test_Packet_LOCAL_INFILE_Data(t *testing.T) {
	var pkt Packet_LOCAL_INFILE_Data
	for _, value := range LOCAL_INFILE_Data_test_packets {
		pkt = Packet_LOCAL_INFILE_Data{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_LOCAL_INFILE_Data_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_LOCAL_INFILE_Data
	for i := 0; i < b.N; i++ {
		pkt = Packet_LOCAL_INFILE_Data{}
		LOCAL_INFILE_Data_test_packets[0].packet.offset = 0
		pkt.FromPacket(context, LOCAL_INFILE_Data_test_packets[0].packet)
	}
}

func Benchmark_Packet_LOCAL_INFILE_Data_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_LOCAL_INFILE_Data{}
	pkt.FromPacket(context, LOCAL_INFILE_Data_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_LOCAL_INFILE_Data_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_LOCAL_INFILE_Data{}
	pkt.FromPacket(context, LOCAL_INFILE_Data_test_packets[0].packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
