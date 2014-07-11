package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildNulTerminatedString(t *testing.T) {
	var values = []struct {
		in   string
		out  []byte
	}{
		{in: "",                out: []byte{0x0,}},
        {in: "A",               out: []byte{0x41, 0x0,}},
        {in: "ABC",             out: []byte{0x41, 0x42, 0x43, 0x0,}},
        {in: "ABCABCABCABC",    out: []byte{0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x0,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildNulTerminatedString(value.in), value.out, "")
	}
}

func Benchmark_BuildNulTerminatedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildNulTerminatedString("ABC")
	}
}

func Test_GetNulTerminatedString(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  string
	}{
		{in: []byte{0x00,}, out: ""},
        {in: []byte{0x41, 0x0}, out: "A"},
        {in: []byte{0x41, 0x42, 0x43, 0x0,}, out: "ABC"},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetNulTerminatedString(), value.out, "")
	}
}

func Benchmark_GetNulTerminatedString(b *testing.B) {
	packet := Packet{ data: []byte{0x41, 0x42, 0x43, 0x0,}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
        packet.offset = 0
		packet.GetNulTerminatedString()
	}
}
