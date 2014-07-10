package MySQLProtocol

func BuildNulTerminatedString(value string) (data []byte) {
    data = make([]byte, len(value)+1)
    copy(data, []byte(value))
    return data
}
