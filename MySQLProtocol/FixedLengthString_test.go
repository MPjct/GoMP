package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildFixedLengthString(t *testing.T) {
	var values = []struct {
		in   string
		size uint
		out  []byte
		eop  bool
	}{
		{in: "", eop: false, size: 0, out: []byte{}},
		{in: "A", eop: false, size: 1, out: []byte{0x41}},
		{in: "ABC", eop: false, size: 3, out: []byte{0x41, 0x42, 0x43}},
		{in: "ABCABCABCABC", eop: false, size: 12, out: []byte{0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43, 0x41, 0x42, 0x43}},
		{in: "A", eop: false, size: 3, out: []byte{0x41, 0x0, 0x0}},
		{in: "ABC", eop: false, size: 1, out: []byte{0x41}},
		{in: "ABC", eop: true, out: []byte{0x41, 0x42, 0x43}},
		{in: "", eop: false, size: 1, out: []byte{0x0}},
	}

	for _, value := range values {
		if value.eop {
			assert.Equal(t, BuildFixedLengthString(value.in), value.out, "")
		} else {
			assert.Equal(t, BuildFixedLengthString(value.in, value.size), value.out, "")
		}
	}
}

func Benchmark_BuildFixedLengthString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthString("ABC", 3)
	}
}

func Test_GetFixedLengthString(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out string
	}{
		{in: []byte{}, out: ""},
		{in: []byte{0x41}, out: "A"},
		{in: []byte{0x41, 0x42, 0x43}, out: "ABC"},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthString(), value.out, "")
	}
    
    values = []struct {
		in  []byte
		out string
	}{
		{in: []byte{0x41}, out: "A"},
		{in: []byte{0x41, 0x42, 0x43}, out: "A"},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthString(1), value.out, "")
	}
}

func Benchmark_GetFixedLengthString(b *testing.B) {
	proto := Proto{data: []byte{0x41, 0x42, 0x43}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.offset = 0
		proto.GetFixedLengthString()
	}
}
