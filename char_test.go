package types

import "fmt"

func ExampleChar_char() {
	c := NewChar("a")
	fmt.Printf("c: %v\n", c)

	// Output:
	// c: a
}

func ExampleChar_string() {
	c := NewChar("hoge")
	fmt.Printf("c: %v\n", c)

	// Output:
	// c: h
}
