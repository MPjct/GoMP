package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Has_Flag(t *testing.T) {
	var values = []struct {
		value  uint64
		flag   uint64
		result bool
	}{
		{value: 0x0, flag: 0x1, result: false},
		{value: 0x1, flag: 0x1, result: true},
	}

	for _, value := range values {
		assert.Equal(t, Has_Flag(value.value, value.flag), value.result, "")
	}
}

func Benchmark_Has_Flag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Has_Flag(0x0, 0x0)
	}
}

func Test_Set_Flag(t *testing.T) {
	var values = []struct {
		value  uint64
		flag   uint64
		result uint64
	}{
		{value: 0x0, flag: 0x1, result: 0x1},
		{value: 0x1, flag: 0x1, result: 0x1},
		{value: 0x1, flag: 0x0, result: 0x1},
	}

	for _, value := range values {
		assert.Equal(t, Set_Flag(value.value, value.flag), value.result, "")
	}
}

func Benchmark_Set_Flag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set_Flag(0x0, 0x0)
	}
}

func Test_Remove_Flag(t *testing.T) {
	var values = []struct {
		value  uint64
		flag   uint64
		result uint64
	}{
		{value: 0x00, flag: 0x01, result: 0x00},
		{value: 0x10, flag: 0x10, result: 0x00},
		{value: 0x11, flag: 0x01, result: 0x10},
	}

	for _, value := range values {
		assert.Equal(t, Remove_Flag(value.value, value.flag), value.result, "")
	}
}

func Benchmark_Remove_Flag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Remove_Flag(0x0, 0x0)
	}
}
