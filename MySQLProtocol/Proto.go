package MySQLProtocol

import "fmt"
import "math"
import "strings"
import "encoding/hex"

type Proto struct {
	data   []byte
	offset uint
}

func (proto *Proto) HasRemainingData() bool {
	return uint(len(proto.data))-proto.offset > 0
}

func DumpPacket(data []byte) {
	fmt.Printf("%s", PacketToString(data))
}

func PacketToString(data []byte) (value string) {
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

func StringToPacket(value string) (data []byte) {
	lines := strings.Split(value, "\n")
	data = make([]byte, 0, 16*len(lines))
    var values []string

	for _, line := range lines {
        if len(line) == 0 {
            continue
        }
		if len(line) < 51 {
			values = strings.Split(line, " ")
		} else {
            values = strings.Split(line[:51], " ")
        }
		for _, val := range values {
			i, _ := hex.DecodeString(val)
			data = append(data, i...)
		}
	}

	return data
}
