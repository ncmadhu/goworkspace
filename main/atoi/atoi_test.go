package atoi

import "testing"

func TestAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"Test0",
			args{"-"},
			0,
		},
		{
			"Test1",
			args{"   -42"},
			-42,
		},
		{
			"Test2",
			args{"42"},
			42,
		},
		{
			"Test3",
			args{"4193 with words"},
			4193,
		},
		{
			"Test4",
			args{"words and 987"},
			0,
		},
		{
			"Test5",
			args{"-91283472332"},
			-2147483648,
		},
		{
			"Test6",
			args{""},
			0,
		},
		{
			"Test7",
			args{"+1"},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atoi(tt.args.str); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
