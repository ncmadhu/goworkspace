package compareversion

import "testing"

func TestCompareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test1",
			args{"0.1", "1.1"},
			-1,
		},
		{
			"Test2",
			args{"1.0.1", "1"},
			1,
		},
		{
			"Test3",
			args{"7.5.2.4", "7.5.3"},
			-1,
		},
		{
			"Test4",
			args{"1.01", "1.001"},
			0,
		},
		{
			"Test5",
			args{"1.0", "1.0.0"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("CompareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
