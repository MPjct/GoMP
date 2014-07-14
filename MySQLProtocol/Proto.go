package MySQLProtocol

import "fmt"
import "math"

type Proto struct {
	data   []byte
	offset uint
}

func (proto *Proto) HasRemainingData() bool {
	return uint(len(proto.data))-proto.offset > 0
}

func DumpPacket(data []byte) {
	fmt.Printf("%s", DumpPacketToString(data))
}

func DumpPacketToString(data []byte) (value string) {
	var i int
	var start int
	var end int

	var first []byte
	var second []byte
	var strval string

	var fstr string
	var sstr string

	value += "\n"

	for i = 0; i < len(data); i += 16 {
		start = i
		end = int(math.Min(float64(len(data)), float64(start+16)))

		if end-start > 8 {
			first = data[start : start+8]
			second = data[start+8 : end]
		} else {
			first = data[start:end]
			second = []byte{}
		}

		strval = ""

		for _, char := range data[start:end] {
			if char >= 32 && char <= 126 {
				strval += fmt.Sprintf("%c", char)
			} else {
				strval += "."
			}
		}

		fstr = fmt.Sprintf("% x", first)
		sstr = fmt.Sprintf("% x", second)

		value += fmt.Sprintf("%-27s%-27s%s\n", fstr, sstr, strval)
	}
	return value
}
