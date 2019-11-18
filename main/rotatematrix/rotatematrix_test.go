package rotatematrix

import "testing"

func TestRotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Test1",
			args{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}},
		},
		{
			"Test2",
			args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.matrix)
		})
	}
}
