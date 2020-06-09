package util

import (
	"encoding/json"
	"testing"
)

func BenchmarkContains(b *testing.B) {
	strs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < b.N; i++ {
		if !Contains(strs, 9) {
			b.Fatal("ToIntSlice error")
		}
	}
}

func BenchmarkContains1(b *testing.B) {
	strs := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < b.N; i++ {
		if !Contains(strs, "9") {
			b.Fatal("ToIntSlice error")
		}
	}
}

func BenchmarkContains2(b *testing.B) {
	strs := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		if !Contains(strs, int64(9)) {
			b.Fatal("ToIntSlice error")
		}
	}
}

func BenchmarkObj2Json(b *testing.B) {
	var req struct {
		MExtendUserInfo map[string]string
		SGuid           string
	}
	req.MExtendUserInfo = make(map[string]string)
	req.SGuid = "1284502258"
	for i := 0; i < b.N; i++ {
		json.MarshalIndent(req, "", "    ")
	}
}

func TestToXXXSlice(t *testing.T) {
	s := "1,2,3,4"
	arr := ToIntSlice(s, ",")
	if len(arr) != 4 {
		t.Fatal("ToIntSlice error")
	}

	arr2 := ToInt32Slice(s, ",")
	if len(arr2) != 4 {
		t.Fatal("ToInt32Slice error")
	}

	arr3 := ToInt64Slice(s, ",")
	if len(arr3) != 4 {
		t.Fatal("ToInt64Slice error")
	}

	arr4 := ToFloat32Slice(s, ",")
	if len(arr4) != 4 {
		t.Fatal("ToFloat32Slice error")
	}

	arr5 := ToFloat64Slice(s, ",")
	if len(arr5) != 4 {
		t.Fatal("ToFloat64Slice error")
	}
}

func TestObj2Json(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"string", args{"abcdef^*"}, `"abcdef^*"`},
		{"html", args{"http://www.baidu.com/?q=abc&n=123<>"}, `"http://www.baidu.com/?q=abc&n=123<>"`},
		{"number", args{123}, `123`},
		{"float", args{123.456}, `123.456`},
		{"struct", args{struct {
			S string `json:"s"`
			N int    `json:"n"`
		}{"abc", 123}}, `{"s":"abc","n":123}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Obj2Json(tt.args.v); got != tt.want {
				t.Errorf("Obj2Json() = %s, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		list   interface{}
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"string", args{[]string{"a", "b", "c"}, "b"}, true},
		{"int", args{[]int{1, 2, 3}, 2}, true},
		{"int8", args{[]int8{1, -2, 3}, -2}, true},
		{"uint8", args{[]uint8{1, 2, 3}, 2}, true},
		{"int16", args{[]int16{1000, -2000, 3000}, -2000}, true},
		{"uint16", args{[]uint16{1000, 2000, 3000}, 2000}, true},
		{"int32", args{[]int32{1000000000, -2000000000, 2100000000}, -2000000000}, true},
		{"uint32", args{[]uint32{1000000000, 2000000000, 4200000000}, 2000000000}, true},
		{"int64", args{[]int64{10000000000000, -20000000000000, 30000000000000}, -20000000000000}, true},
		{"uint64", args{[]uint64{10000000000000, 20000000000000, 30000000000000}, 20000000000000}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
