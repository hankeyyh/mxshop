package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Field = zap.Field

var (
	Object = zap.Object

	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	//String      = zap.String
	//Stringp     = zap.Stringp
	Uint       = zap.Uint
	Uintp      = zap.Uintp
	Uint64     = zap.Uint64
	Uint64p    = zap.Uint64p
	Uint32     = zap.Uint32
	Uint32p    = zap.Uint32p
	Uint16     = zap.Uint16
	Uint16p    = zap.Uint16p
	Uint8      = zap.Uint8
	Uint8p     = zap.Uint8p
	Binary     = zap.Binary
	ByteString = zap.ByteString
	Uintptr    = zap.Uintptr
	Uintptrp   = zap.Uintptrp
	Time       = zap.Time
	Timep      = zap.Timep
	Duration   = zap.Duration
	Durationp  = zap.Durationp
	// Remove avoid global interface conflicts
	// Errorw      = zap.Error
	NamedError = zap.NamedError
	//Stringer   = zap.Stringer

	Stack     = zap.Stack
	StackSkip = zap.StackSkip
)

func Bool(key string, val bool) Field {
	if val {
		return zap.String(key, "true")
	}
	return zap.String(key, "false")
}

func Boolp(key string, val *bool) Field {
	if val == nil {
		return zap.Reflect(key, nil)
	}
	return Bool(key, *val)
}

func String(key string, val string) Field {
	return zap.String(key, truncatedString(val))
}

func Stringp(key string, val *string) Field {
	if val == nil {
		return zap.Stringp(key, val)
	}
	return zap.String(key, truncatedString(*val))
}

const MaxFieldValueLength = 8192

func truncatedString(s string) string {
	if len(s) > MaxFieldValueLength {
		return s[:MaxFieldValueLength] + "...(truncated)"
	}
	return s
}

func Stringer(key string, val fmt.Stringer) Field {
	return zap.Stringer(key, stringerFieldValue{value: val})
}

type stringerFieldValue struct {
	value fmt.Stringer
}

func (f stringerFieldValue) String() string {
	s := f.value.String()
	return truncatedString(s)
}

type anyFieldValue struct {
	value interface{}
}

func (f anyFieldValue) String() string {
	var s string
	e, err := json.Marshal(f.value)
	if err != nil {
		s = fmt.Sprintf("%+v", f.value)
	} else {
		s = string(e)
	}

	return truncatedString(s)
}

func Any(key string, value interface{}) Field {
	switch val := value.(type) {
	case zapcore.ObjectMarshaler:
		return Object(key, val)
	case bool:
		return Bool(key, val)
	case *bool:
		return Boolp(key, val)
	case complex128:
		return Complex128(key, val)
	case *complex128:
		return Complex128p(key, val)
	case complex64:
		return Complex64(key, val)
	case *complex64:
		return Complex64p(key, val)
	case float64:
		return Float64(key, val)
	case *float64:
		return Float64p(key, val)
	case float32:
		return Float32(key, val)
	case *float32:
		return Float32p(key, val)
	case int:
		return Int(key, val)
	case *int:
		return Intp(key, val)
	case int64:
		return Int64(key, val)
	case *int64:
		return Int64p(key, val)
	case int32:
		return Int32(key, val)
	case *int32:
		return Int32p(key, val)
	case int16:
		return Int16(key, val)
	case *int16:
		return Int16p(key, val)
	case int8:
		return Int8(key, val)
	case *int8:
		return Int8p(key, val)
	case string:
		return String(key, val)
	case *string:
		return Stringp(key, val)
	case uint:
		return Uint(key, val)
	case *uint:
		return Uintp(key, val)
	case uint64:
		return Uint64(key, val)
	case *uint64:
		return Uint64p(key, val)
	case uint32:
		return Uint32(key, val)
	case *uint32:
		return Uint32p(key, val)
	case uint16:
		return Uint16(key, val)
	case *uint16:
		return Uint16p(key, val)
	case uint8:
		return Uint8(key, val)
	case *uint8:
		return Uint8p(key, val)
	case uintptr:
		return Uintptr(key, val)
	case *uintptr:
		return Uintptrp(key, val)
	case time.Time:
		return Time(key, val)
	case *time.Time:
		return Timep(key, val)
	case time.Duration:
		return Duration(key, val)
	case *time.Duration:
		return Durationp(key, val)
	case error:
		return NamedError(key, val)
	case fmt.Stringer:
		return Stringer(key, stringerFieldValue{value: val})
	default:
		return Stringer(key, anyFieldValue{value: val})
	}
}
