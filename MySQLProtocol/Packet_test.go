package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var Test_CompressPacket_values = []struct {
	sequence_id uint8
	input       []byte
	output      []byte
}{
	{
		sequence_id: 0,
		input: StringToPacket(`
2e 00 00 00 03 73 65 6c    65 63 74 20 22 30 31 32    .....select "012
33 34 35 36 37 38 39 30    31 32 33 34 35 36 37 38    3456789012345678
39 30 31 32 33 34 35 36    37 38 39 30 31 32 33 34    9012345678901234
35 22                                                 5"
`),
		output: StringToPacket(`
20 00 00 00 32 00 00 d2    63 60 60 60 2e 4e cd 49    ....2...c....N.I
4d 2e 51 50 32 30 34 32    36 31 35 33 b7 b0 c4 cd    M.QP20426153....
52 02 04 00 00 ff ff                                  R......
`),
	},
	{
		sequence_id: 0,
		input: StringToPacket(`
21 00 00 00 03 73 65 6c    65 63 74 20 40 40 76 65    !....select @@ve
72 73 69 6f 6e 5f 63 6f    6d 6d 65 6e 74 20 6c 69    rsion_comment li
6d 69 74 20 31                                        mit 1
`),
		output: StringToPacket(`
25 00 00 00 00 00 00 21    00 00 00 03 73 65 6c 65    %......!....sele
63 74 20 40 40 76 65 72    73 69 6f 6e 5f 63 6f 6d    ct @@version_com
6d 65 6e 74 20 6c 69 6d    69 74 20 31                ment limit 1
`),
	},
}

func Test_CompressPacket(t *testing.T) {
	for _, value := range Test_CompressPacket_values {
		assert.Equal(t, CompressPacket(value.sequence_id, value.input), value.output, "")
	}
}

func Benchmark_CompressPacket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressPacket(Test_CompressPacket_values[0].sequence_id, Test_CompressPacket_values[0].input)
	}
}

func Test_DecompressPacket(t *testing.T) {
	for _, value := range Test_CompressPacket_values {
		in := Proto{data: value.output}
		out := Proto{data: value.input}
		assert.Equal(t, DecompressPacket(in), out, "")
	}
}

func Benchmark_DecompressPacket(b *testing.B) {
	in := Proto{data: Test_CompressPacket_values[0].output}
	for i := 0; i < b.N; i++ {
		in.offset = 0
		DecompressPacket(in)
	}
}
