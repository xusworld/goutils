package goutils

import (
	"reflect"
)

func Compare(lhs, rhs interface{}) int8 {
	if !isComparable(lhs, rhs) {
		return -2
	}

	lhsVal := reflect.ValueOf(lhs)
	rhsVal := reflect.ValueOf(rhs)

	switch lhsVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compareInt(lhsVal, rhsVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return compareUint(lhsVal, rhsVal)
	case reflect.Float32, reflect.Float64:
		return compareFloat(lhsVal, rhsVal)
	default:
		return -2
	}
}

func isComparable(lhs, rhs interface{}) bool {
	lhsVal := reflect.ValueOf(lhs)
	rhsVal := reflect.ValueOf(rhs)
	return lhsVal.Kind() == rhsVal.Kind()
}

func compareInt(lhs, rhs reflect.Value) int8 {
	switch {
	case lhs.Int() < rhs.Int():
		return -1
	case lhs.Int() == rhs.Int():
		return 0
	case lhs.Int() > rhs.Int():
		return 1
	}
	return -2
}

func compareUint(lhs, rhs reflect.Value) int8 {
	switch {
	case lhs.Uint() < rhs.Uint():
		return -1
	case lhs.Uint() == rhs.Uint():
		return 0
	case lhs.Uint() > rhs.Uint():
		return 1
	}
	return -2
}

func compareFloat(lhs, rhs reflect.Value) int8 {
	switch {
	case lhs.Float() < rhs.Float():
		return -1
	case lhs.Float() == rhs.Float():
		return 0
	case lhs.Float() > rhs.Float():
		return 1
	}
	return -2
}

func CompareGT(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) > 0
}

func CompareGE(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) >= 0
}

func CompareLT(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) == -1
}

func CompareLE(lhs, rhs interface{}) bool {
	return Compare(lhs, rhs) == -1 || Compare(lhs, rhs) == 0
}

