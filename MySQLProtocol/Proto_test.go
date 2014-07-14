package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Proto_HasRemainingData(t *testing.T) {
	proto := Proto{data: []byte{0x0}}

	assert.Equal(t, proto.HasRemainingData(), true, "")
	proto.offset = 1
	assert.Equal(t, proto.HasRemainingData(), false, "")
}

func Benchmark_Proto_HasRemainingData(b *testing.B) {
	proto := Proto{data: []byte{0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.HasRemainingData()
	}
}

func Test_Proto_PacketToString(t *testing.T) {

	var values = []struct {
		data []byte
		str  string
	}{
		{
			data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
			str:  "\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n",
		},
		{
			data: []byte{0x05, 0x00, 0x00, 0x05, 0xfe, 0x00, 0x00, 0x02, 0x00},
			str:  "\n05 00 00 05 fe 00 00 02    00                         .........\n",
		},
	}

	for _, value := range values {
		assert.Equal(t, PacketToString(value.data), value.str, "")
	}
}

func Benchmark_Proto_PacketToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PacketToString([]byte{0x0})
	}
}

func Test_Proto_StringToPacket(t *testing.T) {

	var values = []struct {
		data []byte
		str  string
	}{
		{
			data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
			str:  "\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n",
		},
		{
			data: []byte{0x05, 0x00, 0x00, 0x05, 0xfe, 0x00, 0x00, 0x02, 0x00},
			str:  "\n05 00 00 05 fe 00 00 02 00\n",
		},
	}

	for _, value := range values {
		assert.Equal(t, StringToPacket(value.str), value.data, "")
	}
}

func Benchmark_StringToPacket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToPacket("\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n")
	}
}

func Test_Proto_StringToPacketToString(t *testing.T) {

    var values = []string {
        "\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n",
        "\n05 00 00 05 fe 00 00 02    00                         .........\n",
    }

	for _, value := range values {
		assert.Equal(t, PacketToString(StringToPacket(value)), value, "")
	}
}

func Benchmark_StringToPacketToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PacketToString(StringToPacket("\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n"))
	}
}
