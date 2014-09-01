package MySQLProtocol

type Field_TIME struct {
    is_negative uint8
    days uint32
    hours uint8
    minutes uint8
    seconds uint8
    micro_seconds uint32
}

func (field Field_TIME) Build() []byte {
    length := field.Size()
    data := make([]byte, 0, length)
    
    data = append(data, byte(length))
    
    if length >= 8 {
        data = append(data, BuildFixedLengthInteger1(field.is_negative)...)
        data = append(data, BuildFixedLengthInteger4(field.days)...)
        data = append(data, BuildFixedLengthInteger1(field.hours)...)
        data = append(data, BuildFixedLengthInteger1(field.minutes)...)
        data = append(data, BuildFixedLengthInteger1(field.seconds)...)
    }
    
    if length == 12 {
        data = append(data, BuildFixedLengthInteger4(field.micro_seconds)...)
    }
    
	return data
}

func (field *Field_TIME) Get(proto *Proto) {
	length := proto.GetFixedLengthInteger1()
    
    if length >= 8 {
        field.is_negative = proto.GetFixedLengthInteger1()
        field.days = proto.GetFixedLengthInteger4()
        field.hours = proto.GetFixedLengthInteger1()
        field.minutes = proto.GetFixedLengthInteger1()
        field.seconds = proto.GetFixedLengthInteger1()
    }
    
    if length == 12 {
        field.micro_seconds = proto.GetFixedLengthInteger4()
    }
}

func (field Field_TIME) GetType() byte {
	return MYSQL_TYPE_TIME
}

func (field Field_TIME) Size() uint64 {
    if field.micro_seconds != 0 {
        return 12
    }
    
    if field.days != 0 || field.hours != 0 || field.minutes != 0 || field.seconds != 0 {
        return 8
    }
    
    return 0
}
