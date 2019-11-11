package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test0",
			args{[]int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}},
			[][]int{{-5, 1, 4}, {-5, 2, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -2, 4}, {-2, -1, 3}, {-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"Test1",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
