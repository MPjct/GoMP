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
