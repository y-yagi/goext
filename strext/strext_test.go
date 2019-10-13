package strext_test

import (
	"testing"

	"github.com/y-yagi/goext/strext"
)

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: " ", want: true},
		{input: "  ", want: true},
		{input: "\t", want: true},
		{input: "a", want: false},
		{input: "　", want: false},
	}

	for _, test := range tests {
		result := strext.IsBlank(test.input)

		if test.want != result {
			t.Errorf("expected %v, but %v. input: %v", test.want, result, test.input)
		}
	}
}
