package goutils

import (
	"errors"
	"fmt"
	"reflect"
)

func Contains(slice interface{}, target interface{}) bool {
	sliceVal := reflect.ValueOf(slice)
	kind := sliceVal.Kind()

	if kind == reflect.Slice || kind == reflect.Array {

		for idx := 0; idx < sliceVal.Len(); idx++ {
			if sliceVal.Index(idx).Interface() == target {
				return true
			}
		}
	}

	return false
}


// Fill fills elements of array with value
func Fill(in interface{}, fillValue interface{}) (interface{}, error) {
	inValue := reflect.ValueOf(in)
	inKind := inValue.Type().Kind()
	if inKind != reflect.Slice && inKind != reflect.Array {
		return nil, errors.New(" Can only fill slices and arrays")
	}

	inType := reflect.TypeOf(in).Elem()
	value := reflect.ValueOf(fillValue)
	if inType != value.Type() {
		return nil, fmt.Errorf(
			" Cannot fill '%s' with '%s'", reflect.TypeOf(in), value.Type(),
		)
	}

	length := inValue.Len()
	newSlice := reflect.SliceOf(reflect.TypeOf(fillValue))
	in = reflect.MakeSlice(newSlice, length, length).Interface()
	inValue = reflect.ValueOf(in)

	for i := 0; i < length; i++ {
		inValue.Index(i).Set(value)
	}
	return in, nil
}