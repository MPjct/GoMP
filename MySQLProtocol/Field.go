package MySQLProtocol

type Field_interface interface {
	Build() []byte
	Get(proto *Proto)
	GetType() byte
	Size() uint64
}

func (proto *Proto) GetField(field_type byte) (field Field_interface) {
	switch field_type {
	case MYSQL_TYPE_DECIMAL:
		field := new(Field_DECIMAL)
		field.Get(proto)
		return field
	case MYSQL_TYPE_TINY:
		field := new(Field_TINY)
		field.Get(proto)
		return field
	case MYSQL_TYPE_SHORT:
		field := new(Field_SHORT)
		field.Get(proto)
		return field
	case MYSQL_TYPE_LONG:
		field := new(Field_LONG)
		field.Get(proto)
		return field
	case MYSQL_TYPE_FLOAT:
		field := new(Field_FLOAT)
		field.Get(proto)
		return field
	case MYSQL_TYPE_DOUBLE:
		field := new(Field_DOUBLE)
		field.Get(proto)
		return field
	case MYSQL_TYPE_NULL:
		field := new(Field_NULL)
		field.Get(proto)
		return field
	case MYSQL_TYPE_TIMESTAMP:
		field := new(Field_TIMESTAMP)
		field.Get(proto)
		return field
	case MYSQL_TYPE_LONGLONG:
		field := new(Field_LONGLONG)
		field.Get(proto)
		return field
	case MYSQL_TYPE_INT24:
		field := new(Field_INT24)
		field.Get(proto)
		return field
	case MYSQL_TYPE_DATE:
		field := new(Field_DATE)
		field.Get(proto)
		return field
	case MYSQL_TYPE_TIME:
		field := new(Field_TIME)
		field.Get(proto)
		return field
	case MYSQL_TYPE_DATETIME:
		field := new(Field_DATETIME)
		field.Get(proto)
		return field
	case MYSQL_TYPE_YEAR:
		field := new(Field_YEAR)
		field.Get(proto)
		return field
	case MYSQL_TYPE_NEWDATE:
		field := new(Field_NEWDATE)
		field.Get(proto)
		return field
	case MYSQL_TYPE_VARCHAR:
		field := new(Field_VARCHAR)
		field.Get(proto)
		return field
	case MYSQL_TYPE_BIT:
		field := new(Field_BIT)
		field.Get(proto)
		return field
	case MYSQL_TYPE_TIMESTAMP2:
		field := new(Field_TIMESTAMP2)
		field.Get(proto)
		return field
	case MYSQL_TYPE_DATETIME2:
		field := new(Field_DATETIME2)
		field.Get(proto)
		return field
	case MYSQL_TYPE_TIME2:
		field := new(Field_TIME2)
		field.Get(proto)
		return field
	case MYSQL_TYPE_NEWDECIMAL:
		field := new(Field_NEWDECIMAL)
		field.Get(proto)
		return field
	case MYSQL_TYPE_ENUM:
		field := new(Field_ENUM)
		field.Get(proto)
		return field
	case MYSQL_TYPE_SET:
		field := new(Field_SET)
		field.Get(proto)
		return field
	case MYSQL_TYPE_TINY_BLOB:
		field := new(Field_TINY_BLOB)
		field.Get(proto)
		return field
	case MYSQL_TYPE_MEDIUM_BLOB:
		field := new(Field_MEDIUM_BLOB)
		field.Get(proto)
		return field
	case MYSQL_TYPE_LONG_BLOB:
		field := new(Field_LONG_BLOB)
		field.Get(proto)
		return field
	case MYSQL_TYPE_BLOB:
		field := new(Field_BLOB)
		field.Get(proto)
		return field
	case MYSQL_TYPE_VAR_STRING:
		field := new(Field_VAR_STRING)
		field.Get(proto)
		return field
	case MYSQL_TYPE_STRING:
		field := new(Field_STRING)
		field.Get(proto)
		return field
	case MYSQL_TYPE_GEOMETRY:
		field := new(Field_GEOMETRY)
		field.Get(proto)
		return field
	}
    return nil
}
