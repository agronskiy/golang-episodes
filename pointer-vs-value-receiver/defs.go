package main

import (
	"fmt"
	"reflect"
	"unsafe"
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

// Value type receivers
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

// Check that interface holds copy of value
func demonstrateCopyingByInterface() {
	var (
		v  T                 = T{i: 0} // Note zero here
		iv ValueMethodCaller = v
	)

	v.i = 10 // Changing the original object

	fmt.Printf("Original value: \t%#v\n", v)
	fmt.Printf("Interface value: \t%#v\n", reflect.ValueOf(iv))
}

func demonstrateDifferentTypesInTheInterface() {
	var (
		x int32   = 0
		y float32 = 10.0
	)

	var iface interface{} = x
	fmt.Printf("Interface value: \t%#v\n", reflect.ValueOf(iface))

	// Trying to take address of the value. This will not compile, but imagine it were here
	// var px *int32 = &reflect.ValueOf(iface)

	iface = y
	fmt.Printf("Interface value: \t%#v\n", reflect.ValueOf(iface))
}

func demonstrateCopyingInTheInterface() {
	var iface interface{} = (int32)(0)

	// This takes address of the value. Unsafe but works. Not guaranteed to work
	// after possible implementation change!
	var px uintptr = (*[2]uintptr)(unsafe.Pointer(&iface))[1]

	iface = (int32)(1)

	var py uintptr = (*[2]uintptr)(unsafe.Pointer(&iface))[1]
	fmt.Printf("First pointer %#v,  second pointer %#v", px, py)

}
