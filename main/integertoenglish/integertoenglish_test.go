package integertoenglish

import "testing"

func TestNumberToWords(t *testing.T) {
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
			args{1234567891},
			"One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
		{
			"Test2",
			args{1234567},
			"One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			"Test3",
			args{12345},
			"Twelve Thousand Three Hundred Forty Five",
		},
		{
			"Test4",
			args{123},
			"One Hundred Twenty Three",
		},
		{
			"Test5",
			args{1000010},
			"One Million Ten",
		},
		{
			"Test6",
			args{0},
			"Zero",
		},
		{
			"Test7",
			args{2},
			"Two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberToWords(tt.args.num); got != tt.want {
				t.Errorf("NumberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
