package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

var NULL_Bitmap_test_packets = []struct {
	packet Proto
}{
	{packet: Proto{data: StringToPacket(`
00 00
`)}},
}

func Test_NULL_Bitmap(t *testing.T) {
	for _, value := range NULL_Bitmap_test_packets {

		null_bitmap := value.packet.GetNullBitmap(9, 2)

		assert.Equal(t, null_bitmap.GetNullBitmapSize(), uint64(2), "")

		assert.Equal(t, null_bitmap.BuildNullBitmap(), []byte{0x0, 0x0}, "")

		null_bitmap.SetField(8, true)

		assert.Equal(t, null_bitmap.BuildNullBitmap(), []byte{0x0, 0x04}, "")

		null_bitmap.SetField(8, false)

		assert.Equal(t, null_bitmap.BuildNullBitmap(), []byte{0x0, 0x00}, "")
	}
}
