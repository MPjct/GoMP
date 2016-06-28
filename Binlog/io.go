package Binlog

type Binlog_IO struct {
	err            error
	mode_stream    bool
	mode_write     bool
	binlog_version uint16
}

//func ReadBinlog(name string) Binlog_IO {}

//func StreamBinlogs(name string) Binlog_IO {}

//func WriteBinlog(name string) Binlog_IO {}
