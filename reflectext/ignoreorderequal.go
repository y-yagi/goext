package reflectext

import (
	"reflect"
	"sort"
)

// IgnoreOrderEqual reports where x and y are equal. This ignores order in Array.
func IgnoreOrderEqual(x, y interface{}) bool {
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}

	if v1.Kind() != reflect.Slice {
		return reflect.DeepEqual(x, y)
	}

	if v1.Len() == 0 {
		return reflect.DeepEqual(x, y)
	}

	v := v1.Index(0)
	switch v.Kind() {
	case reflect.String:
		sort.Strings(x.([]string))
		sort.Strings(y.([]string))
	case reflect.Int:
		sort.Ints(x.([]int))
		sort.Ints(y.([]int))
	}

	return reflect.DeepEqual(x, y)
}
