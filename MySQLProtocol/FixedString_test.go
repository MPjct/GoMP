package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildFixedString(t *testing.T) {
	var values = []struct {
		in   string
		out  []byte
	}{
		{in: "",                out: []byte{}},
        {in: "A",               out: []byte{0x41,}},
        {in: "ABC",             out: []byte{0x41, 0x42, 0x43,}},
        {in: "ABCABCABCABC",    out: []byte{0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedString(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedString("ABC")
	}
}
