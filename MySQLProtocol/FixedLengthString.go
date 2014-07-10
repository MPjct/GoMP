package MySQLProtocol

func BuildFixedLengthString(value string, size ...uint) (data []byte) {
    var datasize uint
    if len(size) == 0 {
        datasize = uint(len(value))
        data = make([]byte, datasize) 
    } else if size[0] > uint(len(value)) {
        data = make([]byte, size[0]) 
        datasize = uint(len(value))
    } else {
        data = make([]byte, size[0]) 
        datasize = size[0]
    }
    copy(data, []byte(value[:datasize]))
    return data
}
