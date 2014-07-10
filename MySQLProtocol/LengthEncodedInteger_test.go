package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildLengthEncodedInt(t *testing.T) {
	var values = []struct {
		in   uint64
		out  []byte
	}{
		{in: 0,        out: []byte{0x00,}},
        {in: 251,      out: []byte{0xFC, 0xFB, 0x00, }},
        {in: 252,      out: []byte{0xFC, 0xFC, 0x00, }},
        {in: 65536,    out: []byte{0xFD, 0x00, 0x00, 0x01 }},
        {in: 16777216, out: []byte{0xFE, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00 }},
        {in: 33554432, out: []byte{0xFE, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00 }},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildLengthEncodedInt(value.in), value.out, "")
	}
}

func Benchmark_BuildLengthEncodedInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildLengthEncodedInt(uint64(i))
	}
}

func Test_GetLengthEncodedInt(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint64
	}{
		{in: []byte{0xFA}, out: 250},
        {in: []byte{0xFC, 0xFB, 0x00}, out: 251},
        {in: []byte{0xFD, 0xFB, 0x00, 0x10}, out: 1048827},
        {in: []byte{0xFE, 0xFB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10}, out: 1152921504606847227},
        
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetLengthEncodedInt(), value.out, "")
	}
}

func Benchmark_GetLengthEncodedInt(b *testing.B) {
	packet := Packet{ data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetLengthEncodedInt()
	}
}
