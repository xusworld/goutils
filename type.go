package goutils


func InterfaceSliceToInt(interfaceSlice []interface{}) []int64 {
	int64Slice := make([]int64, len(interfaceSlice))
	for idx, inter := range interfaceSlice {
		interfaceSlice[idx] = inter.(int64)
	}
	return int64Slice
}

func InterfaceSliceToUint(interfaceSlice []interface{}) []uint64 {
	uint64Slice := make([]uint64, len(interfaceSlice))
	for idx, inter := range interfaceSlice {
		interfaceSlice[idx] = inter.(uint64)
	}
	return uint64Slice
}
