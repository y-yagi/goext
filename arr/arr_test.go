package arr_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/y-yagi/goext/arr"
)

func TestContains(t *testing.T) {
	value := "today"
	list := []string{"today", "week", "month", "all"}

	if !arr.Contains(list, value) {
		t.Errorf("Expect contains returns true but false. list: %v, value: '%s'", list, value)
	}

	value = "tomorrow"
	if arr.Contains(list, value) {
		t.Errorf("Expect contains returns false but true. list: %v, value: '%s'", list, value)
	}
}

func TestJoin(t *testing.T) {
	list := []string{"a", "b", "c"}
	expected := "a-b-c"

	result := arr.Join(list, "-")
	if result != expected {
		t.Errorf("expected %v, but %v.", expected, result)
	}
}

func TestUniqueStrings(t *testing.T) {
	list := []string{"abc", "bcd", "cde", "abc", "def"}
	expected := []string{"abc", "bcd", "cde", "def"}

	result := arr.UniqueStrings(list)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("UniqueStrings() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkContains_found(b *testing.B) {
	value := "today"
	list := []string{"today", "week", "month", "all"}
	for n := 0; n < b.N; n++ {
		arr.Contains(list, value)
	}
}

func BenchmarkContains_notfound(b *testing.B) {
	value := "unuexist"
	list := []string{"today", "week", "month", "all"}

	for n := 0; n < b.N; n++ {
		arr.Contains(list, value)
	}
}
