package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Packet_HasRemainingData(t *testing.T) {
    packet := Packet{ data: []byte{0x0}}
    
    assert.Equal(t, packet.HasRemainingData(), true, "")
    packet.offset = 1
    assert.Equal(t, packet.HasRemainingData(), false, "")
}

func Benchmark_Packet_HasRemainingData(b *testing.B) {
    packet := Packet{ data: []byte{0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.HasRemainingData()
	}
}
