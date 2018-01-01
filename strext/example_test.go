package strext_test

import (
	"fmt"

	"github.com/y-yagi/goext/strext"
)

func ExampleIsBlank() {
	fmt.Println(strext.IsBlank(""))
	// Output: true
}

func ExampleIsBlank_mutiplespaces() {
	fmt.Println(strext.IsBlank(" "))
	// Output: true
}

func ExampleIsBlank_multibytespace() {
	fmt.Println(strext.IsBlank("　"))
	// Output: false
}
