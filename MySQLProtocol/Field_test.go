package MySQLProtocol

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

var fields_test_packets = []struct {
	field_type byte
	packet     Proto
	context    Context
}{
	{field_type: MYSQL_TYPE_STRING,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_VARCHAR,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_VAR_STRING,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_ENUM,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_SET,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_LONG_BLOB,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_MEDIUM_BLOB,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_BLOB,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TINY_BLOB,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_GEOMETRY,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_BIT,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_DECIMAL,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_NEWDECIMAL,
		packet:  Proto{data: StringToPacket(`03 66 6f 6f`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_LONGLONG,
		packet:  Proto{data: StringToPacket(`01 00 00 00 00 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_LONG,
		packet:  Proto{data: StringToPacket(`01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_INT24,
		packet:  Proto{data: StringToPacket(`01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_SHORT,
		packet:  Proto{data: StringToPacket(`01 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_YEAR,
		packet:  Proto{data: StringToPacket(`01 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TINY,
		packet:  Proto{data: StringToPacket(`01`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_DOUBLE,
		packet:  Proto{data: StringToPacket(`66 66 66 66 66 66 24 40`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_FLOAT,
		packet:  Proto{data: StringToPacket(`33 33 23 41`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_DATE,
		packet:  Proto{data: StringToPacket(`04 da 07 0a 11`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_NEWDATE,
		packet:  Proto{data: StringToPacket(`04 da 07 0a 11`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_DATETIME,
		packet:  Proto{data: StringToPacket(`0b da 07 0a 11 13 1b 1e 01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_DATETIME2,
		packet:  Proto{data: StringToPacket(`0b da 07 0a 11 13 1b 1e 01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIMESTAMP,
		packet:  Proto{data: StringToPacket(`0b da 07 0a 11 13 1b 1e 01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIMESTAMP2,
		packet:  Proto{data: StringToPacket(`0b da 07 0a 11 13 1b 1e 01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIME,
		packet:  Proto{data: StringToPacket(`0c 01 78 00 00 00 13 1b 1e 01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIME,
		packet:  Proto{data: StringToPacket(`08 01 78 00 00 00 13 1b 1e`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIME,
		packet:  Proto{data: StringToPacket(`01`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIME2,
		packet:  Proto{data: StringToPacket(`0c 01 78 00 00 00 13 1b 1e 01 00 00 00`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIME2,
		packet:  Proto{data: StringToPacket(`08 01 78 00 00 00 13 1b 1e`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_TIME2,
		packet:  Proto{data: StringToPacket(`01`)},
		context: Context{}},

	{field_type: MYSQL_TYPE_NULL,
		packet:  Proto{data: StringToPacket(``)},
		context: Context{}},
}

func Test_Packet_fields(t *testing.T) {
	var field Field_interface
	var packet Proto
	var error_msg string
	var expected_size int
	var size int

	for _, value := range fields_test_packets {
		field = value.packet.GetField(value.field_type)

		error_msg = fmt.Sprintf("Error with type %d", value.field_type)

		assert.Equal(t, field.GetType(), value.field_type, error_msg)
		assert.Equal(t, field.Build(), value.packet.data, error_msg)

		expected_size = int(field.PacketSize())
		size = len(value.packet.data)

		error_msg = fmt.Sprintf("Error with type %d, size %d expected %d", value.field_type, expected_size, size)

		assert.Equal(t, expected_size, size, error_msg)
	}

	// Test unknown fields
	packet = Proto{data: StringToPacket(``)}
	field = packet.GetField(MYSQL_TYPE_UNKNOWN)
	assert.Equal(t, field, nil, "Unknown field does not return nil")
}
