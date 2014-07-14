package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Proto_HasRemainingData(t *testing.T) {
	proto := Proto{data: []byte{0x0}}

	assert.Equal(t, proto.HasRemainingData(), true, "")
	proto.offset = 1
	assert.Equal(t, proto.HasRemainingData(), false, "")
}

func Benchmark_Proto_HasRemainingData(b *testing.B) {
	proto := Proto{data: []byte{0x0}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.HasRemainingData()
	}
}

func Test_Proto_StringToPacket(t *testing.T) {

	var values = []struct {
		data []byte
		str  string
	}{
		{
			data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
			str:  "\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n",
		},
		{
			data: []byte{0x05, 0x00, 0x00, 0x05, 0xfe, 0x00, 0x00, 0x02, 0x00},
			str:  "\n05 00 00 05 fe 00 00 02 00\n",
		},
		{
			data: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
				0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
				0x1F, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29,
				0x2A, 0x2B, 0x2C, 0x2D, 0x2E, 0x2F, 0x30, 0x31, 0x32, 0x33, 0x34,
				0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E, 0x3F,
				0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A,
				0x4B, 0x4C, 0x4D, 0x4E, 0x4F, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55,
				0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F, 0x60,
				0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B,
				0x6C, 0x6D, 0x6E, 0x6F, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76,
				0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F, 0x80, 0x81,
				0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x8A, 0x8B, 0x8C,
				0x8D, 0x8E, 0x8F, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
				0x98, 0x99, 0x9A, 0x9B, 0x9C, 0x9D, 0x9E, 0x9F, 0xA0, 0xA1, 0xA2,
				0xA3, 0xA4, 0xA5, 0xA6, 0xA7, 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD,
				0xAE, 0xAF, 0xB0, 0xB1, 0xB2, 0xB3, 0xB4, 0xB5, 0xB6, 0xB7, 0xB8,
				0xB9, 0xBA, 0xBB, 0xBC, 0xBD, 0xBE, 0xBF, 0xC0, 0xC1, 0xC2, 0xC3,
				0xC4, 0xC5, 0xC6, 0xC7, 0xC8, 0xC9, 0xCA, 0xCB, 0xCC, 0xCD, 0xCE,
				0xCF, 0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9,
				0xDA, 0xDB, 0xDC, 0xDD, 0xDE, 0xDF, 0xE0, 0xE1, 0xE2, 0xE3, 0xE4,
				0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF,
				0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9, 0xFA,
				0xFB, 0xFC, 0xFD, 0xFE, 0xFF},
			str: `
00 01 02 03 04 05 06 07    08 09 0a 0b 0c 0d 0e 0f    ................
10 11 12 13 14 15 16 17    18 19 1a 1b 1c 1d 1e 1f    ................
20 21 22 23 24 25 26 27    28 29 2a 2b 2c 2d 2e 2f     !"#$%&'()*+,-./
30 31 32 33 34 35 36 37    38 39 3a 3b 3c 3d 3e 3f    0123456789:;<=>?
40 41 42 43 44 45 46 47    48 49 4a 4b 4c 4d 4e 4f    @ABCDEFGHIJKLMNO
50 51 52 53 54 55 56 57    58 59 5a 5b 5c 5d 5e 5f    PQRSTUVWXYZ[\]^_
60 61 62 63 64 65 66 67    68 69 6a 6b 6c 6d 6e 6f    ` + "`" + `abcdefghijklmno
70 71 72 73 74 75 76 77    78 79 7a 7b 7c 7d 7e 7f    pqrstuvwxyz{|}~.
80 81 82 83 84 85 86 87    88 89 8a 8b 8c 8d 8e 8f    ................
90 91 92 93 94 95 96 97    98 99 9a 9b 9c 9d 9e 9f    ................
a0 a1 a2 a3 a4 a5 a6 a7    a8 a9 aa ab ac ad ae af    ................
b0 b1 b2 b3 b4 b5 b6 b7    b8 b9 ba bb bc bd be bf    ................
c0 c1 c2 c3 c4 c5 c6 c7    c8 c9 ca cb cc cd ce cf    ................
d0 d1 d2 d3 d4 d5 d6 d7    d8 d9 da db dc dd de df    ................
e0 e1 e2 e3 e4 e5 e6 e7    e8 e9 ea eb ec ed ee ef    ................
f0 f1 f2 f3 f4 f5 f6 f7    f8 f9 fa fb fc fd fe ff    ................
`,
		},
		{
			data: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
			str:  "\n01 02 03 04 05 06 07 08  09 10 11 12 13 14 15 16   |................|\n",
		},
	}

	for _, value := range values {
		DumpPacket(value.data)
		assert.Equal(t, StringToPacket(value.str), value.data, "")
	}
}

func Benchmark_StringToPacket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToPacket("\n01 02 03 04 05 06 07 08    09 10 11 12 13 14 15 16    ................\n")
	}
}
