package MySQLProtocol

type Field_interface interface {
	Build() []byte
	Get(proto *Proto)
	GetType() byte
	Size() uint64
}

func GetField(field_type uint64, proto *Proto) (field Field_interface) {
	switch field_type {
	case MYSQL_TYPE_TINY:
		return nil
	case MYSQL_TYPE_SHORT:
		return nil
	case MYSQL_TYPE_LONG:
		return nil
	case MYSQL_TYPE_FLOAT:
		return nil
	case MYSQL_TYPE_DOUBLE:
		return nil
	case MYSQL_TYPE_NULL:
		return nil
	case MYSQL_TYPE_TIMESTAMP:
		return nil
	case MYSQL_TYPE_LONGLONG:
		return nil
	case MYSQL_TYPE_INT24:
		return nil
	case MYSQL_TYPE_DATE:
		return nil
	case MYSQL_TYPE_TIME:
		return nil
	case MYSQL_TYPE_DATETIME:
		return nil
	case MYSQL_TYPE_YEAR:
		return nil
	case MYSQL_TYPE_NEWDATE:
		return nil
	case MYSQL_TYPE_TIMESTAMP2:
		return nil
	case MYSQL_TYPE_DATETIME2:
		return nil
	case MYSQL_TYPE_TIME2:
		return nil
	case MYSQL_TYPE_STRING, MYSQL_TYPE_VARCHAR, MYSQL_TYPE_VAR_STRING,
		MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_LONG_BLOB,
		MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_BLOB, MYSQL_TYPE_TINY_BLOB,
		MYSQL_TYPE_GEOMETRY, MYSQL_TYPE_BIT, MYSQL_TYPE_DECIMAL,
		MYSQL_TYPE_NEWDECIMAL:
		field := new(Field_STRING)
		field.Get(proto)
		return field
	default:
		return nil
	}
}
