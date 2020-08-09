package reflectext_test

import (
	"testing"

	"github.com/y-yagi/goext/reflectext"
)

func TestIgnoreOrderEqual(t *testing.T) {
	tests := []struct {
		x    interface{}
		y    interface{}
		want bool
	}{
		{x: []string{"c", "b", "a"}, y: []string{"a", "b", "c"}, want: true},
		{x: []int{3, 2, 1}, y: []int{1, 2, 3}, want: true},
	}

	for _, tt := range tests {
		result := reflectext.IgnoreOrderEqual(tt.x, tt.y)

		if tt.want != result {
			t.Errorf("expected '%v', but '%v'. x: '%v', y: '%v'", tt.want, result, tt.x, tt.y)
		}
	}
}
