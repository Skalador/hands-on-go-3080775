// challenges/types-primitive/begin/main.go
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 12

	// print the value of the local variable "val"
	fmt.Printf("%T, local variable %v\n", val, val)

	// print the value of the package-level variable "val"
	printGlobal()

	// update the package-level variable "val" and print it again
	val = 13
	fmt.Printf("%T, local variable %v\n", val, val)

	// print the pointer address of the local variable "val"
	var a *int
	a = &val
	fmt.Printf("Address of val %d\n", a)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*a = 27
	fmt.Printf("%T, local variable %v\n", val, val)
}

func printGlobal() {
	fmt.Printf("Global Variable %v of Type %T\n", val, val)
}
