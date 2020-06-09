package goutils

import (
	"testing"
)

func TestAnyTypeToString(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Int",
			args: args{
				42,
			},
			want: "42",
		},
		{
			name: "Test Int8",
			args: args{
				int8(42),
			},
			want: "42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyTypeToString(tt.args.val); got != tt.want {
				t.Errorf("AnyTypeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}