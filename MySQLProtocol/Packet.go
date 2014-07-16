package MySQLProtocol

import "compress/zlib"
import "io/ioutil"
import "bytes"

type Packet struct {
	sequence_id uint8
}

func CompressPacket(sequence_id uint8, input []byte) (output []byte) {
	return output
}

func DecompressPacket(input Proto) (output Proto) {
	compressed_payload_length := input.GetFixedLengthInteger3()
	input.GetFixedLengthInteger1()
	uncompressed_payload_length := input.GetFixedLengthInteger3()

	if uncompressed_payload_length == 0 {
		output.data = input.ExtractSlice(uint(compressed_payload_length))
	} else {
		b := bytes.NewReader(input.ExtractSlice(uint(compressed_payload_length)))
		r, err := zlib.NewReader(b)
		if err == nil {
			panic("Decompression failed!")
		}
		data, err := ioutil.ReadAll(r)
		if err == nil {
			panic("Decompression failed!")
		}
		output.data = data
		r.Close()
	}
	return output
}

func EncryptPacket(packets ...Packet) (data []byte) {
	return data
}

func DecryptPacket(data Proto) (packets []Packet) {
	return packets
}
