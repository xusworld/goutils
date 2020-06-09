package goutils

import (
	"errors"
	"math"
	"reflect"
)

func Max(slice []interface{}) interface{} {
	if len(slice) == 0 {
		return -2
	}

	kind := reflect.ValueOf(slice[0]).Kind()

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return MaxIntSlice(InterfaceSliceToInt(slice))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return MaxUintSlice(InterfaceSliceToUint(slice))
	default:
		return nil
	}
}

func MaxIntSlice(slice []int64) int64 {
	if len(slice) == 0 {
		return math.MinInt64
	}

	max := slice[0]

	for idx := 1; idx < len(slice); idx++ {
		if max < slice[idx] {
			max = slice[idx]
		}
	}
	return max
}

func MaxUintSlice(slice []uint64) uint64 {
	if len(slice) == 0 {
		return 0
	}

	max := slice[0]

	for idx := 1; idx < len(slice); idx++ {
		if max < slice[idx] {
			max = slice[idx]
		}
	}
	return max
}

func MinIntSlice(slice []int64) int64 {
	if len(slice) == 0 {
		return math.MaxInt64
	}

	min := slice[0]
	for idx := 1; idx < len(slice); idx++ {
		if min > slice[idx] {
			min = slice[idx]
		}
	}
	return min
}

func MinUintSlice(slice []uint64) uint64 {
	if len(slice) == 0 {
		return math.MaxUint64
	}

	min := slice[0]
	for idx := 1; idx < len(slice); idx++ {
		if min > slice[idx] {
			min = slice[idx]
		}
	}
	return min
}

// TODO Slice/Array 操作
// 1. Slice/Array sum
// 2. Slice/Array product
// 3. Slice/Array max
// 4. Slice/Array min

func SumInts(i interface{}) int64 {
	result, err := sum(i)
	if err != nil {
		return 0
	}
	return result.(int64)
}

func SumIUints(i interface{}) uint64 {
	result, err := sum(i)
	if err != nil {
		return 0
	}
	return result.(uint64)
}

func SumFloats(i interface{}) float64 {
	result, err := sum(i)
	if err != nil {
		return 0.0
	}
	return result.(float64)
}

func sum(i interface{}) (interface{}, error) {
	val := reflect.ValueOf(i)

	if val.Kind() != reflect.Slice || val.Kind() != reflect.Array {
		return val, errors.New("input is not a slice or array")
	}

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var result int64
		for idx := 0; idx < val.Len(); idx++ {
			result += val.Index(idx).Int()
		}
		return result, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var result uint64
		for idx := 0; idx < val.Len(); idx++ {
			result += val.Index(idx).Uint()
		}
		return result, nil
	case reflect.Float32, reflect.Float64:
		var result float64
		for idx := 0; idx < val.Len(); idx++ {
			result += val.Index(idx).Float()
		}
		return result, nil
	}

}
