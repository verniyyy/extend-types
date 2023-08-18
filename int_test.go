package types

import (
	"fmt"
)

func ExampleInt() {
	i := NewInt(6)
	fmt.Printf("i: %v\n", i)

	// Output:
	// i: 6
}
