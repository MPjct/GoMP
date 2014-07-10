package MySQLProtocol

func BuildFixedString(size uint, value string) (data []byte) {
    data = make([]byte, size)
    if size > uint(len(value)) {
        size = uint(len(value))
    }
    copy(data, []byte(value[:size]))
    return data
}
