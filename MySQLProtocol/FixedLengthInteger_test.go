package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_MySQLProtocol_BuildFixedInt1(t *testing.T) {
	var values = []struct {
		in   uint8
		out  [1]byte
	}{
		{in: 0, out: [1]byte{0x00,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedInt1(value.in), value.out, "")
	}
}

func Benchmark_MySQLProtocol_BuildFixedInt_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt1(uint8(i))
	}
}

func Test_MySQLProtocol_BuildFixedInt2(t *testing.T) {
	var values = []struct {
		in   uint16
		out  [2]byte
	}{
		{in: 0, out: [2]byte{0x00, 0x00,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedInt2(value.in), value.out, "")
	}
}

func Benchmark_MySQLProtocol_BuildFixedInt_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt2(uint16(i))
	}
}

func Test_MySQLProtocol_BuildFixedInt3(t *testing.T) {
	var values = []struct {
		in   uint32
		out  [3]byte
	}{
		{in: 0, out: [3]byte{0x00, 0x00, 0x00,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedInt3(value.in), value.out, "")
	}
}

func Benchmark_MySQLProtocol_BuildFixedInt_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt3(uint32(i))
	}
}

func Test_MySQLProtocol_BuildFixedInt4(t *testing.T) {
	var values = []struct {
		in   uint32
		out  [4]byte
	}{
		{in: 0, out: [4]byte{0x00, 0x00, 0x00, 0x00,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedInt4(value.in), value.out, "")
	}
}

func Benchmark_MySQLProtocol_BuildFixedInt_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt4(uint32(i))
	}
}

func Test_MySQLProtocol_BuildFixedInt8(t *testing.T) {
	var values = []struct {
		in   uint64
		out  [8]byte
	}{
		{in: 0, out: [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedInt8(value.in), value.out, "")
	}
}

func Benchmark_MySQLProtocol_BuildFixedInt_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt8(uint64(i))
	}
}
