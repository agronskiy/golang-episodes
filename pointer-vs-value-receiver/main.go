package main

import (
	"fmt"
)

// PointerMethodCaller is the testing interface
type PointerMethodCaller interface {
	pointerMethod()
}

// ValueMethodCaller is the testing interface
type ValueMethodCaller interface {
	valueMethod()
}

// T is the testing type
type T struct {
	i int8
}

// Pointer type receiver
func (receiver *T) pointerMethod() {
	fmt.Printf("Pointer method called on \t%#v with address %p\n", *receiver, receiver)
}

// Value type receiver
func (receiver T) valueMethod() {
	fmt.Printf("Value method called on \t%#v with address %p\n", receiver, &receiver)
}

// Calling methods on interfaces
func callValueMethodOnInterface(v ValueMethodCaller) {
	v.valueMethod()
}

func callPointerMethodOnInterface(p PointerMethodCaller) {
	p.pointerMethod()
}

func main() {
	var (
		val     T  = T{}
		pointer *T = &val
	)

	fmt.Printf("Value created: %#v with address %p\n", val, &val)
	fmt.Printf("Pointer created for the object \t%#v with address %p\n", *pointer, pointer)

	val.valueMethod()
	pointer.pointerMethod()

	// Cross-calling on different receivers
	fmt.Println("### Cross-calling on different receivers ###")
	val.pointerMethod()
	pointer.valueMethod()

	// Interface part
	fmt.Println("### Interface part ###")
	callValueMethodOnInterface(val)
	callPointerMethodOnInterface(pointer)

	callValueMethodOnInterface(pointer)
	callPointerMethodOnInterface(val)
}
