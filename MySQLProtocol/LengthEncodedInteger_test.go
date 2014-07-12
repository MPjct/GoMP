package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildLengthEncodedInteger(t *testing.T) {
	var values = []struct {
		in  uint64
		out []byte
	}{
		{in: 0, out: []byte{0x00}},
		{in: 251, out: []byte{0xFC, 0xFB, 0x00}},
		{in: 252, out: []byte{0xFC, 0xFC, 0x00}},
		{in: 65536, out: []byte{0xFD, 0x00, 0x00, 0x01}},
		{in: 16777216, out: []byte{0xFE, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00}},
		{in: 33554432, out: []byte{0xFE, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildLengthEncodedInteger(value.in), value.out, "")
	}
}

func Benchmark_BuildLengthEncodedInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildLengthEncodedInteger(uint64(i))
	}
}

func Test_GetLengthEncodedInteger(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint64
	}{
		{in: []byte{0xFA}, out: 250},
		{in: []byte{0xFC, 0xFB, 0x00}, out: 251},
		{in: []byte{0xFD, 0xFB, 0x00, 0x10}, out: 1048827},
		{in: []byte{0xFE, 0xFB, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10}, out: 1152921504606847227},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetLengthEncodedInteger(), value.out, "")
	}
}

func Benchmark_GetLengthEncodedInteger(b *testing.B) {
	proto := Proto{data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
        proto.offset = 0
		proto.GetLengthEncodedInteger()
	}
}

func Test_GetLengthEncodedIntegerSize(t *testing.T) {
	var values = []struct {
		in  uint64
		out uint64
	}{
        {in: 0, out: 1},
		{in: 250, out: 1},
        {in: 0xFFFE, out: 3},
        {in: 0xFFFFFE, out: 4},
        {in: 0xFFFFFF, out: 9},
	}

	for _, value := range values {
		assert.Equal(t, GetLengthEncodedIntegerSize(value.in), value.out, "")
	}
}

func Benchmark_GetLengthEncodedIntegerSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLengthEncodedIntegerSize(0)
	}
}

