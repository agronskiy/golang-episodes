package main

import (
	"fmt"

	"github.com/fatih/color"
)

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
	color.Green("\n### Cross-calling on different receivers ###")
	fmt.Println()
	val.pointerMethod()
	pointer.valueMethod()

	// Interface part
	color.Green("\n### Interface part ###")
	fmt.Println()
	callValueMethodOnInterface(val)
	callPointerMethodOnInterface(pointer)

	callValueMethodOnInterface(pointer)

	// THIS IS WHAT CRASHES (see blog post)
	// callPointerMethodOnInterface(val)

	color.Green("\n### Demonstrate copying in the interface ###")
	fmt.Println()
	demonstrateCopyingByInterface()

	color.Green("\n### Dynamic type in the interface ###")
	fmt.Println()
	demonstrateDifferentTypesInTheInterface()

	color.Green("\n### Always new copy in the interface ###")
	fmt.Println()
	demonstrateCopyingInTheInterface()
}
