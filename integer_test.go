package types

import (
	"fmt"
)

func ExampleInt() {
	i := NewInteger(6)
	fmt.Printf("i: %v\n", i)

	// Output:
	// i: 6
}

func ExampleInt_UInt8() {
	i := NewInteger(9223372036854775807)
	ui8 := i.UInt8()
	fmt.Printf("ui8: %v\n", ui8)

	// Output:
	// ui8: 255
}
