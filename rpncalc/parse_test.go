package rpncalc

import (
	"slices"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"3 2 +", []string{"3", "2", "+"}},
		{"123.456 72 * 44.7 /", []string{"123.456", "72", "*", "44.7", "/"}},
		{"3 2+", []string{"3", "2", "+"}},
		{"123.456 72*44.7/", []string{"123.456", "72", "*", "44.7", "/"}},
	}

	for _, test := range tests {
		have := tokenize(test.input)
		if slices.Compare(test.want, have) != 0 {
			t.Errorf("Wanted %v but got %v", test.want, have)
		}

	}

}
