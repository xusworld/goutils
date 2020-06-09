package goutils

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		slice  interface{}
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Integer Slice True Case",
			args: args{
				slice:  []int{24, 42},
				target: 42,
			},
			want: true,
		},
		{
			name: "Integer Slice False Case",
			args: args{
				slice:  []int{24, 42},
				target: 11,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}