package main

import (
	"fmt"
)

// T is the testing type
type T struct{}

// Pointer type receiver
func (receiver *T) pointerMethod() {
	fmt.Printf("Pointer method on \t%#v with address %p\n", *receiver, receiver)
}

// Value type receiver
func (receiver T) valueMethod() {
	fmt.Printf("Value method on \t%#v with address %p\n", receiver, &receiver)
}

func main() {
	var (
		val     T  = T{}
		pointer *T = &val
	)

	fmt.Printf("Value created \t\t%#v with address %p\n", val, &val)
	fmt.Printf("Pointer created on \t%#v with address %p\n", *pointer, pointer)

	val.valueMethod()
	pointer.pointerMethod()
}
