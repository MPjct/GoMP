package MySQLProtocol

import "compress/zlib"
import "io/ioutil"
import "bytes"

type Packet struct {
	sequence_id uint8
}

func CompressPacket(sequence_id uint8, input []byte) (output []byte) {
	var compressed_payload_length uint32
	var compressed_payload []byte
	var uncompressed_payload_length uint32
	uncompressed_payload_length = uint32(len(input))

	if uncompressed_payload_length < MIN_COMPRESS_LENGTH {
		compressed_payload_length = uncompressed_payload_length
		uncompressed_payload_length = 0
		compressed_payload = input
	} else {
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write(input)
		w.Close()
		compressed_payload = b.Bytes()
		compressed_payload_length = uint32(len(compressed_payload))
	}

	output = make([]byte, compressed_payload_length+7)
	output = append(output, BuildFixedLengthInteger3(compressed_payload_length)...)
	output = append(output, BuildFixedLengthInteger1(sequence_id)...)
	output = append(output, BuildFixedLengthInteger3(uncompressed_payload_length)...)
	output = append(output, compressed_payload...)
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

func EncryptPacket(sequence_id uint8, input []byte) (output []byte) {
	return output
}

func DecryptPacket(input Proto) (output Proto) {
	return output
}
