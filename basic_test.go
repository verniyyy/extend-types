package types

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	b := newBasic(16)
	fmt.Println(b)
}

func TestBasic_PrimitiveTypeName(t *testing.T) {
	b := newBasic(16)
	fmt.Printf("b.PrimitiveTypeName(): %v\n", b.PrimitiveTypeName())
}
