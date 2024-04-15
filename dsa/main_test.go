package main

import "testing"

//func Test_findLongestSubStr(t *testing.T) {
//	type args struct {
//		input string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantRes string
//	}{
//		{
//			name: "",
//			args: args{
//				input: "abc12a",
//			},
//			wantRes: "abc12",
//		},
//		{
//			name: "",
//			args: args{
//				input: "1223444",
//			},
//			wantRes: "234",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotRes := findLongestSubStr(tt.args.input); gotRes != tt.wantRes {
//				t.Errorf("findLongestSubStr() = %v, want %v", gotRes, tt.wantRes)
//			}
//		})
//	}
//}

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				num: "123912",
				k:   1,
			},
			want: "12312",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
