package test

import "testing"

func TestCompareVersion(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "大于",
			args: args{
				lhs: "1.8.6",
				rhs: "1.8.5",
			},
			want: 1,
		},
		{
			name: "大于",
			args: args{
				lhs: "1.8.6.1",
				rhs: "1.8.6",
			},
			want: 1,
		},
		{
			name: "等于",
			args: args{
				lhs: "1.8.5",
				rhs: "1.8.5",
			},
			want: 0,
		},
		{
			name: "小于",
			args: args{
				lhs: "1.8.4",
				rhs: "1.8.5",
			},
			want: -1,
		},
		{
			name: "小于",
			args: args{
				lhs: "1.8.4",
				rhs: "1.8.4.1",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersion(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareVersionGE(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "大于",
			args: args{
				lhs: "1.8.1",
				rhs: "1.8.0",
			},
			want: true,
		},
		{
			name: "大于等于",
			args: args{
				lhs: "1.8.1",
				rhs: "1.8.1",
			},
			want: true,
		},
		{
			name: "大于等于情况为false",
			args: args{
				lhs: "1.8.0",
				rhs: "1.8.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersionGE(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareVersionGE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareVersionGT(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "大于",
			args: args{
				lhs: "1.8.1",
				rhs: "1.8.0",
			},
			want: true,
		},
		{
			name: "大于",
			args: args{
				lhs: "1.8.1",
				rhs: "1.8.1",
			},
			want: false,
		},
		{
			name: "大于等于情况为false",
			args: args{
				lhs: "1.8.0",
				rhs: "1.8.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersionGT(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareVersionGT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareVersionLE(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "小于",
			args: args{
				lhs: "1.8.0",
				rhs: "1.8.1",
			},
			want: true,
		},
		{
			name: "小于等于",
			args: args{
				lhs: "1.8.1",
				rhs: "1.8.1",
			},
			want: true,
		},
		{
			name: "小于等于情况为false",
			args: args{
				lhs: "1.8.2.4",
				rhs: "1.8.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersionLE(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareVersionLE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareVersionLT(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "小于",
			args: args{
				lhs: "1.8.0",
				rhs: "1.8.1",
			},
			want: true,
		},
		{
			name: "小于等于",
			args: args{
				lhs: "1.8.1",
				rhs: "1.8.1",
			},
			want: false,
		},
		{
			name: "小于等于情况为false",
			args: args{
				lhs: "1.8.2.4",
				rhs: "1.8.1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersionLT(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("CompareVersionLT() = %v, want %v", got, tt.want)
			}
		})
	}
}