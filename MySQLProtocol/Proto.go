package MySQLProtocol

import "fmt"
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
	fmt.Printf("\n%s\n", hex.Dump(data))
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
