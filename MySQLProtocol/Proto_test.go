package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Proto_HasRemainingData(t *testing.T) {
	packet := Proto{data: []byte{0x0}}

	assert.Equal(t, packet.HasRemainingData(), true, "")
	packet.offset = 1
	assert.Equal(t, packet.HasRemainingData(), false, "")
}

func Benchmark_Proto_HasRemainingData(b *testing.B) {
	packet := Proto{data: []byte{0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.HasRemainingData()
	}
}
