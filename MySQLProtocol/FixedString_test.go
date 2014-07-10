package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildFixedString(t *testing.T) {
	var values = []struct {
		in   string
        size uint
		out  []byte
	}{
		{in: "",             size: 0,  out: []byte{}},
        {in: "A",            size: 1,  out: []byte{0x41,}},
        {in: "ABC",          size: 3,  out: []byte{0x41, 0x42, 0x43,}},
        {in: "ABCABCABCABC", size: 12, out: []byte{0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43,}},
        {in: "A",            size: 3,  out: []byte{0x41, 0x0, 0x0,}},
        {in: "ABC",          size: 1,  out: []byte{0x41,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedString(value.size, value.in), value.out, "")
	}
}

func Benchmark_BuildFixedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedString(3, "ABC")
	}
}
