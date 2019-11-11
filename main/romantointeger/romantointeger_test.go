package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test1",
			args{"III"},
			3,
		},
		{
			"Test2",
			args{"IV"},
			4,
		},
		{
			"Test3",
			args{"IX"},
			9,
		},
		{
			"Test4",
			args{"LVIII"},
			58,
		},
		{
			"Test5",
			args{"MCMXCIV"},
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
