package types

import (
	"fmt"
)

func ExampleBasic() {
	b := newBasic("foo")
	fmt.Printf("b: %v\n", b)

	// Output:
	// b: foo
}

func ExampleBasic_PrimitiveTypeName_string() {
	b := newBasic("foo")
	fmt.Printf("b.PrimitiveTypeName(): %v\n", b.CoreTypeName())

	// Output:
	// b.PrimitiveTypeName(): string
}

func ExampleBasic_Print() {
	b := newBasic("foo")
	b.Print("b")

	// Output:
	// b: foo
}
