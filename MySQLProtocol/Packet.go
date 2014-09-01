package MySQLProtocol

import "compress/flate"
import "bytes"
import "fmt"

type Packet_interface interface {
    GetPacketSize(Context) uint64
    ToPacket(Context) []byte
    FromPacket(Context, Proto)
}

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
		w, _ := flate.NewWriter(&b, 9)
		w.Write(input)
		w.Close()
		compressed_payload = b.Bytes()
		compressed_payload_length = uint32(len(compressed_payload))
	}

	output = make([]byte, 0, compressed_payload_length+7)
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
		compressedData := input.ExtractSlice(uint(compressed_payload_length))
		b := bytes.NewReader(compressedData)
		r := flate.NewReader(b)
		data := make([]byte, uncompressed_payload_length)
		n, err := r.Read(data)
		if err != nil {
			panic(err.Error())
		}
		if uint32(n) != uncompressed_payload_length {
			DumpPacket(data)
			fmt.Printf("Decompression was too short: %d, expected:%d\n", n, uncompressed_payload_length)
			panic("Decompression was too short!")
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
