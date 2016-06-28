package Binlog

// https://dev.mysql.com/doc/internals/en/binlog-version.html
const Binlog_Version_1 uint16 = 1
const Binlog_Version_2 uint16 = 2
const Binlog_Version_3 uint16 = 3
const Binlog_Version_4 uint16 = 4

// https://dev.mysql.com/doc/internals/en/binlog-event-flag.html
const LOG_EVENT_BINLOG_IN_USE_F uint16 = 0x0001
const LOG_EVENT_FORCED_ROTATE_F uint16 = 0x0002
const LOG_EVENT_THREAD_SPECIFIC_F uint16 = 0x0004
const LOG_EVENT_SUPPRESS_USE_F uint16 = 0x0008
const LOG_EVENT_UPDATE_TABLE_MAP_VERSION_F uint16 = 0x0010
const LOG_EVENT_ARTIFICIAL_F uint16 = 0x0020
const LOG_EVENT_RELAY_LOG_F uint16 = 0x0040
const LOG_EVENT_IGNORABLE_F uint16 = 0x0080
const LOG_EVENT_NO_FILTER_F uint16 = 0x0100
const LOG_EVENT_MTS_ISOLATE_F uint16 = 0x0200
