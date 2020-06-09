package goutils

import (
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		lhs interface{}
		rhs interface{}
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "整型比较-小于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: -1,
		},
		{
			name: "整型比较-等于",
			args: args {
				lhs: 42,
				rhs: 42,
			},
			want: 0,
		},
		{
			name: "整型比较-大于",
			args: args {
				lhs: 43,
				rhs: 42,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareGE(t *testing.T) {
	type args struct {
		lhs interface{}
		rhs interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "整型比较-小于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: false,
		},
		{
			name: "整型比较-等于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: true,
		},
		{
			name: "整型比较-大于",
			args: args {
				lhs: 4,
				rhs: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareGE(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareGT(t *testing.T) {
	type args struct {
		lhs interface{}
		rhs interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "整型比较-小于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: false,
		},
		{
			name: "整型比较-等于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: false,
		},
		{
			name: "整型比较-大于",
			args: args {
				lhs: 4,
				rhs: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareGT(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareGT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareLE(t *testing.T) {
	type args struct {
		lhs interface{}
		rhs interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "整型比较-小于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: true,
		},
		{
			name: "整型比较-等于",
			args: args {
				lhs: 3,
				rhs: 3,
			},
			want: true,
		},
		{
			name: "整型比较-大于",
			args: args {
				lhs: 4,
				rhs: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLE(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareLE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareLT(t *testing.T) {
	type args struct {
		lhs interface{}
		rhs interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "整型比较-小于",
			args: args {
				lhs: 3,
				rhs: 4,
			},
			want: true,
		},
		{
			name: "整型比较-等于",
			args: args {
				lhs: 3,
				rhs: 3,
			},
			want: false,
		},
		{
			name: "整型比较-大于",
			args: args {
				lhs: 4,
				rhs: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLT(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareLT() = %v, want %v", got, tt.want)
			}
		})
	}
}