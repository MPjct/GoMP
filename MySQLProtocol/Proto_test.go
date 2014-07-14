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

func Test_Proto_DumpPacketToString(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
    0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16,
    }
    str := "\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n"

	assert.Equal(t, DumpPacketToString(data), str, "")
}

func Benchmark_Proto_DumpPacketToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DumpPacketToString([]byte{0x0})
	}
}

