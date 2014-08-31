package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"

func Test_NULL_Bitmap(t *testing.T) {
	var null_bitmap NULL_bitmap
    
    null_bitmap.offset = 2
    null_bitmap.num_fields = 9
    null_bitmap.nulls = make([]byte, 2, 2)
    
    assert.Equal(t, null_bitmap.GetNullBitmapSize(), 2, "")
    
    assert.Equal(t, null_bitmap.BuildNullBitmap(), []byte{0x0, 0x0}, "")
    
    null_bitmap.SetField(8, true)
    
    assert.Equal(t, null_bitmap.BuildNullBitmap(), []byte{0x0, 0x04}, "")
    
    null_bitmap.SetField(8, false)
    
    assert.Equal(t, null_bitmap.BuildNullBitmap(), []byte{0x0, 0x00}, "")
}
