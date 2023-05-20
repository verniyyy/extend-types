package types

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	i := NewInt(6)
	i.Print()
}

func TestInt_String(t *testing.T) {
	i := NewInt(6)
	fmt.Printf("i: %v\n", i)
}
