package types

import "fmt"

func ExampleFloat() {
	f := NewFloat(5.5)
	fmt.Printf("f: %v\n", f)

	// Output:
	// f: 5.5
}
