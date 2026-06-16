package keyboardrow

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		words []string
		want  []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
		{[]string{"omg", "type", "row"}, []string{"type", "row"}},
		{[]string{"flag", "sky"}, []string{"flag"}},
		{[]string{"Alaska", "Alaska"}, []string{"Alaska", "Alaska"}},
		{[]string{"Aa"}, []string{"Aa"}},
	}
	for _, tt := range tests {
		got := findWords(tt.words)
		if len(got) == 0 && len(tt.want) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("findWords(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}
}
