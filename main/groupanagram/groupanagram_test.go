package groupanagram

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"Test0",
			args{[]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}},
			[][]string{{"ill"}, {"bar"}, {"max"}, {"doc"}, {"cab"}, {"tin"}, {"pew"}, {"duh"}, {"may"}, {"buy"}},
		},
		{
			"Test1",
			args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
