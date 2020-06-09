package goutils

import (
	"reflect"
	"strconv"
)

func NumToString(i interface{}) string {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(val.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64)
	default:
		return ""
	}
}

func AtoI(str string) int {
	r, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0
	}
	return int(r)
}
func AtoUI(str string) uint {
	r, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0
	}
	return uint(r)
}
func AtoI32(str string) int32 {
	r, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(r)
}
func AtoUI32(str string) uint32 {
	r, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(r)
}
func AtoI64(str string) int64 {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return int64(res)
}
func AtoUI64(str string) uint64 {
	res, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(res)
}
func AtoF32(str string) float32 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return float32(i)
}
func AtoF64(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return res
}
