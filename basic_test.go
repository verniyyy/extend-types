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

func ExampleBasic_PrimitiveTypeName_Foo() {
	type Foo struct{}
	b := newBasic(Foo{})
	fmt.Printf("b.PrimitiveTypeName(): %v\n", b.CoreTypeName())

	// Output:
	// b.PrimitiveTypeName(): Foo
}

func ExampleBasic_Printf() {
	b := newBasic("foo")
	b.Printf("b: %v")

	// Output:
	// b: foo
}

func ExampleBasic_Sprintf() {
	b := newBasic("foo")
	s := b.Sprintf("b: %v")
	fmt.Printf("s: %v\n", s)

	// Output:
	// s: b: foo
}
