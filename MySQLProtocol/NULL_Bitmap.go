package MySQLProtocol


type NULL_bitmap struct {
    num_fields uint16
    offset uint8
    nulls []byte
}

func (null_bitmap NULL_bitmap) BuildNullBitmap() (val []byte) {
	return null_bitmap.nulls
}

func (proto *Proto) GetNullBitmap(num_fields uint16, offset uint8) (null_bitmap NULL_bitmap) {
    null_bitmap.num_fields = num_fields
    null_bitmap.offset = offset
    
    size := null_bitmap.GetNullBitmapSize()
    end := uint64(proto.offset) + size
    
    for uint64(proto.offset) < end {
        val := proto.GetFixedLengthInteger1()
        null_bitmap.nulls = append(null_bitmap.nulls, val)
	}

	return null_bitmap
}

func (null_bitmap NULL_bitmap) GetNullBitmapSize() uint64 {
	return uint64((null_bitmap.num_fields + 7 + uint16(null_bitmap.offset)) / 8)
}

func (null_bitmap NULL_bitmap) SetField(field_pos uint16, null bool) {
    field_index := field_pos + uint16(null_bitmap.offset)
    byte_pos := field_index / 8
    bit_pos  := field_index % 8
    
    if null {
        null_bitmap.nulls[byte_pos] = null_bitmap.nulls[byte_pos] | (1 << bit_pos)
    } else {
        null_bitmap.nulls[byte_pos] = null_bitmap.nulls[byte_pos] &^ (1 << bit_pos)
    }
}
