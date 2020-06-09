package goutils

import (
	"reflect"
	"strconv"
)

// convert any type to string
func AnyTypeToString(val interface{}) string {
	if val == nil {
		return ""
	}
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'E',-1, 64)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.String:
		return value.String()
	case reflect.Ptr:
		return value.Elem().String()
	default:
		return value.Type().String()
	}
}
