package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_AuthSwitchRequest(t *testing.T) {
	var values = []struct {
		packet  Proto
		context Context
	}{
		{packet: Proto{data: StringToPacket(`
2c 00 00 02 fe 6d 79 73    71 6c 5f 6e 61 74 69 76    ,....mysql_nativ
65 5f 70 61 73 73 77 6f    72 64 00 7a 51 67 34 69    e_password.zQg4i
36 6f 4e 79 36 3d 72 48    4e 2f 3e 2d 62 29 41 00    6oNy6=rHN/>-b)A.
`)}, context: Context{}},
	}
	var pkt Packet_AuthSwitchRequest

	for _, value := range values {
		pkt = Packet_AuthSwitchRequest{}
		pkt.FromPacket(value.context, value.packet)
		assert.Equal(t, pkt.ToPacket(value.context), value.packet.data, "")
	}
}

func Benchmark_Packet_AuthSwitchRequest_FromPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	var pkt Packet_AuthSwitchRequest
	var packet = Proto{data: StringToPacket(`
2c 00 00 02 fe 6d 79 73    71 6c 5f 6e 61 74 69 76    ,....mysql_nativ
65 5f 70 61 73 73 77 6f    72 64 00 7a 51 67 34 69    e_password.zQg4i
36 6f 4e 79 36 3d 72 48    4e 2f 3e 2d 62 29 41 00    6oNy6=rHN/>-b)A.
`)}
	for i := 0; i < b.N; i++ {
		pkt = Packet_AuthSwitchRequest{}
		packet.offset = 0
		pkt.FromPacket(context, packet)
	}
}

func Benchmark_Packet_AuthSwitchRequest_GetPacketSize(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_AuthSwitchRequest{}
	var packet = Proto{data: StringToPacket(`
2c 00 00 02 fe 6d 79 73    71 6c 5f 6e 61 74 69 76    ,....mysql_nativ
65 5f 70 61 73 73 77 6f    72 64 00 7a 51 67 34 69    e_password.zQg4i
36 6f 4e 79 36 3d 72 48    4e 2f 3e 2d 62 29 41 00    6oNy6=rHN/>-b)A.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.GetPacketSize(context)
	}
}

func Benchmark_Packet_AuthSwitchRequest_ToPacket(b *testing.B) {
	context := Context{capability: CLIENT_PROTOCOL_41}
	pkt := Packet_AuthSwitchRequest{}
	var packet = Proto{data: StringToPacket(`
2c 00 00 02 fe 6d 79 73    71 6c 5f 6e 61 74 69 76    ,....mysql_nativ
65 5f 70 61 73 73 77 6f    72 64 00 7a 51 67 34 69    e_password.zQg4i
36 6f 4e 79 36 3d 72 48    4e 2f 3e 2d 62 29 41 00    6oNy6=rHN/>-b)A.
`)}
	pkt.FromPacket(context, packet)
	for i := 0; i < b.N; i++ {
		pkt.ToPacket(context)
	}
}
