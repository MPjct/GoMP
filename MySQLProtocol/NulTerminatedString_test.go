package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildNullString(t *testing.T) {
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
        assert.Equal(t, BuildNullString(value.in), value.out, "")
	}
}

func Benchmark_BuildNullString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildNullString("ABC")
	}
}
