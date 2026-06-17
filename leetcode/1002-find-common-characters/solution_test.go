package findcommoncharacters

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"ab", "cd"}, []string{}},
		{[]string{"a", "a"}, []string{"a"}},
		{[]string{"aa", "aa"}, []string{"a", "a"}},
		{[]string{"a", "aa", "aaa"}, []string{"a"}},
		{
			[]string{"acabcddd",
				"bcbdbcbd",
				"baddbadb",
				"cbdddcac",
				"aacbcccd",
				"ccccddda",
				"cababaab",
				"addcaccd"},
			[]string{},
		},
	}

	for _, tt := range tests {
		got := commonChars(tt.words)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("commonChars(%q) = %q, want %q", tt.words, got, tt.want)
		}
	}
}
