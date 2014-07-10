package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_BuildFixedInt1(t *testing.T) {
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

func Benchmark_BuildFixedInt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt1(uint8(i))
	}
}

func Test_GetFixedInt1(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint8
	}{
		{in: []byte{0x00,}, out: 0},
        {in: []byte{0x01,}, out: 1},
        {in: []byte{0xFF,}, out: 255},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetFixedInt1(), value.out, "")
	}
}

func Benchmark_GetFixedInt1(b *testing.B) {
	packet := Packet{ data: []byte{0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetFixedInt1()
	}
}

func Test_BuildFixedInt2(t *testing.T) {
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

func Benchmark_BuildFixedInt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt2(uint16(i))
	}
}

func Test_GetFixedInt2(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint16
	}{
		{in: []byte{0x00, 0x00,}, out: 0},
        {in: []byte{0x01, 0x00,}, out: 1},
        {in: []byte{0xFF, 0x00,}, out: 255},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetFixedInt2(), value.out, "")
	}
}

func Benchmark_GetFixedInt2(b *testing.B) {
	packet := Packet{ data: []byte{0x0, 0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetFixedInt2()
	}
}

func Test_BuildFixedInt3(t *testing.T) {
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

func Benchmark_BuildFixedInt3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt3(uint32(i))
	}
}

func Test_GetFixedInt3(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint32
	}{
		{in: []byte{0x00, 0x00, 0x00,}, out: 0},
        {in: []byte{0x01, 0x00, 0x00,}, out: 1},
        {in: []byte{0xFF, 0x00, 0x00,}, out: 255},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetFixedInt3(), value.out, "")
	}
}

func Benchmark_GetFixedInt3(b *testing.B) {
	packet := Packet{ data: []byte{0x0, 0x0, 0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetFixedInt3()
	}
}

func Test_BuildFixedInt4(t *testing.T) {
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

func Benchmark_BuildFixedInt4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt4(uint32(i))
	}
}

func Test_GetFixedInt4(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint32
	}{
		{in: []byte{0x00, 0x00, 0x00, 0x00,}, out: 0},
        {in: []byte{0x01, 0x00, 0x00, 0x00,}, out: 1},
        {in: []byte{0xFF, 0x00, 0x00, 0x00,}, out: 255},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetFixedInt4(), value.out, "")
	}
}

func Benchmark_GetFixedInt4(b *testing.B) {
	packet := Packet{ data: []byte{0x0, 0x0, 0x0, 0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetFixedInt4()
	}
}

func Test_BuildFixedInt6(t *testing.T) {
	var values = []struct {
		in   uint64
		out  [6]byte
	}{
		{in: 0, out: [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}},
	}
    
    for _, value := range values {
        assert.Equal(t, BuildFixedInt6(value.in), value.out, "")
	}
}

func Benchmark_BuildFixedInt6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt6(uint64(i))
	}
}

func Test_GetFixedInt6(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint64
	}{
		{in: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}, out: 0},
        {in: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00,}, out: 1},
        {in: []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00,}, out: 255},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetFixedInt6(), value.out, "")
	}
}

func Benchmark_GetFixedInt6(b *testing.B) {
	packet := Packet{ data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0,}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetFixedInt6()
	}
}

func Test_BuildFixedInt8(t *testing.T) {
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

func Benchmark_BuildFixedInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildFixedInt8(uint64(i))
	}
}

func Test_GetFixedInt8(t *testing.T) {
    var packet Packet
	var values = []struct {
		in   []byte
		out  uint64
	}{
		{in: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}, out: 0},
        {in: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}, out: 1},
        {in: []byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}, out: 255},
	}
    
    for _, value := range values {
        packet = Packet{ data: value.in}
        assert.Equal(t, packet.GetFixedInt8(), value.out, "")
	}
}

func Benchmark_GetFixedInt8(b *testing.B) {
	packet := Packet{ data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
    b.ResetTimer()
	for i := 0; i < b.N; i++ {
		packet.GetFixedInt8()
	}
}
