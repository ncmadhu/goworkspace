package stringreverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{"Hello World", "dlroW olleH"},
		{"Malayalam", "malayalaM"},
		{"", ""},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.output {
			t.Errorf("Reverse(%q) != %q", test.input, result)
		}
	}
}
