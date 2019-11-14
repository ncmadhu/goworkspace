package strstr

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test0",
			args{"mississippi", "issi"},
			1,
		},
		{
			"Test1",
			args{"hello", "ll"},
			2,
		},
		{
			"Test2",
			args{"a", "a"},
			0,
		},
		{
			"Test3",
			args{"mississippi", "issip"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
