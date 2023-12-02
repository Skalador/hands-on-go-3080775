// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type generic interface {
	int | bool | string
}

// Part 1: print function refactoring
func printAny[T generic](x T) { fmt.Println(x) }

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// Part 2 sum function refactoring
type numbers interface {
	constraints.Integer | constraints.Float
}

// sum sums a slice of any type
func sum[T numbers](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	// call non-generic print functions
	//printString("Hello")
	//printInt(42)
	//printBool(true)

	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call sum function
	//fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	fmt.Println("result", sum(1, 2, 3))
	fmt.Println("result", sum(1.0, 2.0, 3.3))
}
