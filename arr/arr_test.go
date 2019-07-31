package arr

import "testing"

func TestContains(t *testing.T) {
	value := "today"
	list := []string{"today", "week", "month", "all"}

	if !Contains(list, value) {
		t.Errorf("Expect contains returns true but false. list: %v, value: '%s'", list, value)
	}

	value = "tomorrow"
	if Contains(list, value) {
		t.Errorf("Expect contains returns false but true. list: %v, value: '%s'", list, value)
	}
}

func TestJoin(t *testing.T) {
	list := []string{"a", "b", "c"}
	expected := "a-b-c"

	result := Join(list, "-")
	if result != expected {
		t.Errorf("expected %v, but %v.", expected, result)
	}
}
