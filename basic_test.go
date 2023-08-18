package types

import (
	"fmt"
)

func ExampleBasic() {
	b := newBasic(16)
	fmt.Printf("b: %v\n", b)

	// Output:
	// b: 16
}

func ExampleBasic_PrimitiveTypeName() {
	b := newBasic(16)
	fmt.Printf("b.PrimitiveTypeName(): %v\n", b.PrimitiveTypeName())

	// Output:
	// b.PrimitiveTypeName(): int
}
