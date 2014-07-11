package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildFixedLengthInteger1(t *testing.T) {
	var values = []struct {
		in  uint8
		out [1]byte
	}{
		{in: 0, out: [1]byte{0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildFixedLengthInteger1(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedLengthInteger1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthInteger1(uint8(i))
	}
}

func Test_GetFixedLengthInteger1(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint8
	}{
		{in: []byte{0x00}, out: 0},
		{in: []byte{0x01}, out: 1},
		{in: []byte{0xFF}, out: 255},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthInteger1(), value.out, "")
	}
}

func Benchmark_GetFixedLengthInteger1(b *testing.B) {
	proto := Proto{data: []byte{0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.GetFixedLengthInteger1()
	}
}

func Test_BuildFixedLengthInteger2(t *testing.T) {
	var values = []struct {
		in  uint16
		out [2]byte
	}{
		{in: 0, out: [2]byte{0x00, 0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildFixedLengthInteger2(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedLengthInteger2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthInteger2(uint16(i))
	}
}

func Test_GetFixedLengthInteger2(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint16
	}{
		{in: []byte{0x00, 0x00}, out: 0},
		{in: []byte{0x01, 0x00}, out: 1},
		{in: []byte{0xFF, 0x00}, out: 255},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthInteger2(), value.out, "")
	}
}

func Benchmark_GetFixedLengthInteger2(b *testing.B) {
	proto := Proto{data: []byte{0x0, 0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.GetFixedLengthInteger2()
	}
}

func Test_BuildFixedLengthInteger3(t *testing.T) {
	var values = []struct {
		in  uint32
		out [3]byte
	}{
		{in: 0, out: [3]byte{0x00, 0x00, 0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildFixedLengthInteger3(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedLengthInteger3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthInteger3(uint32(i))
	}
}

func Test_GetFixedLengthInteger3(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint32
	}{
		{in: []byte{0x00, 0x00, 0x00}, out: 0},
		{in: []byte{0x01, 0x00, 0x00}, out: 1},
		{in: []byte{0xFF, 0x00, 0x00}, out: 255},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthInteger3(), value.out, "")
	}
}

func Benchmark_GetFixedLengthInteger3(b *testing.B) {
	proto := Proto{data: []byte{0x0, 0x0, 0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.GetFixedLengthInteger3()
	}
}

func Test_BuildFixedLengthInteger4(t *testing.T) {
	var values = []struct {
		in  uint32
		out [4]byte
	}{
		{in: 0, out: [4]byte{0x00, 0x00, 0x00, 0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildFixedLengthInteger4(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedLengthInteger4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthInteger4(uint32(i))
	}
}

func Test_GetFixedLengthInteger4(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint32
	}{
		{in: []byte{0x00, 0x00, 0x00, 0x00}, out: 0},
		{in: []byte{0x01, 0x00, 0x00, 0x00}, out: 1},
		{in: []byte{0xFF, 0x00, 0x00, 0x00}, out: 255},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthInteger4(), value.out, "")
	}
}

func Benchmark_GetFixedLengthInteger4(b *testing.B) {
	proto := Proto{data: []byte{0x0, 0x0, 0x0, 0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.GetFixedLengthInteger4()
	}
}

func Test_BuildFixedLengthInteger6(t *testing.T) {
	var values = []struct {
		in  uint64
		out [6]byte
	}{
		{in: 0, out: [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildFixedLengthInteger6(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedLengthInteger6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthInteger6(uint64(i))
	}
}

func Test_GetFixedLengthInteger6(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint64
	}{
		{in: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, out: 0},
		{in: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00}, out: 1},
		{in: []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00}, out: 255},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthInteger6(), value.out, "")
	}
}

func Benchmark_GetFixedLengthInteger6(b *testing.B) {
	proto := Proto{data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.GetFixedLengthInteger6()
	}
}

func Test_BuildFixedLengthInteger8(t *testing.T) {
	var values = []struct {
		in  uint64
		out [8]byte
	}{
		{in: 0, out: [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, value := range values {
		assert.Equal(t, BuildFixedLengthInteger8(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedLengthInteger8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedLengthInteger8(uint64(i))
	}
}

func Test_GetFixedLengthInteger8(t *testing.T) {
	var proto Proto
	var values = []struct {
		in  []byte
		out uint64
	}{
		{in: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, out: 0},
		{in: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, out: 1},
		{in: []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, out: 255},
	}

	for _, value := range values {
		proto = Proto{data: value.in}
		assert.Equal(t, proto.GetFixedLengthInteger8(), value.out, "")
	}
}

func Benchmark_GetFixedLengthInteger8(b *testing.B) {
	proto := Proto{data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.GetFixedLengthInteger8()
	}
}
