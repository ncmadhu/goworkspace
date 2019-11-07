package romannumerals

import "testing"

func TestIntToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test1",
			args{3},
			"III",
		},
		{
			"Test2",
			args{4},
			"IV",
		},
		{
			"Test3",
			args{9},
			"IX",
		},
		{
			"Test4",
			args{58},
			"LVIII",
		},
		{
			"Test5",
			args{1994},
			"MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToRoman(tt.args.num); got != tt.want {
				t.Errorf("IntToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
