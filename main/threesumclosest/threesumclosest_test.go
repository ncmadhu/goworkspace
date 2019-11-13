package threesumclosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				[]int{-1, 2, 1, -4},
				1,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("ThreeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
