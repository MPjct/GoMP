package MySQLProtocol

type Attributes struct {
	attributes [][2]string
}

func (attributes *Attributes) SetAttribute(attribute string, value string) bool {
	if attributes.attributes == nil {
		attributes.attributes = make([][2]string, 0, 1)
	}
	if attributes.GetAttribute(attribute) == nil {
		attributes.attributes = append(attributes.attributes, [2]string{attribute, value})
	} else {
		for index, attr := range attributes.attributes {
			if attr[0] == attribute {
				attributes.attributes[index] = [2]string{attribute, value}
			}
		}
	}
	return true
}

func (attributes *Attributes) RemoveAttribute(attribute string) bool {
	if attributes.attributes == nil {
		return false
	}
	if attributes.GetAttribute(attribute) == nil {
		return false
	}
	for index, attr := range attributes.attributes {
		if attr[0] == attribute {
			copy(attributes.attributes[index:], attributes.attributes[index+1:])
			attributes.attributes = attributes.attributes[:len(attributes.attributes)-1]
		}
	}
	return true
}

func (attributes *Attributes) GetAttribute(attribute string) (value *string) {
	if attributes.attributes == nil {
		return nil
	}
	for _, attr := range attributes.attributes {
		if attr[0] == attribute {
			return &attr[1]
		}
	}
	return nil
}

func (attributes Attributes) BuildAttributes() (val []byte) {
	if attributes.attributes == nil {
		return val
	}
	var totalattributesize uint64
	for _, attr := range attributes.attributes {
		totalattributesize += GetLengthEncodedStringSize(attr[0])
		totalattributesize += GetLengthEncodedStringSize(attr[1])
	}
	val = append(val, BuildLengthEncodedInteger(totalattributesize)...)
	for _, attr := range attributes.attributes {
		val = append(val, BuildLengthEncodedString(attr[0])...)
		val = append(val, BuildLengthEncodedString(attr[1])...)
	}
	return val
}

func (proto *Proto) GetAttributes() (attributes Attributes) {
	attributes = Attributes{}

	datasize := proto.GetLengthEncodedInteger()
	end := uint64(proto.offset) + datasize

	for uint64(proto.offset) < end {
		key := proto.GetLengthEncodedString()
		val := proto.GetLengthEncodedString()
		attributes.SetAttribute(key, val)
	}
	return attributes
}

func (attributes Attributes) GetAttributesSize() uint64 {
	return uint64(len(attributes.BuildAttributes()))
}
