package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// := is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "short" in this case.
	f := "short"
	fmt.Println(f)
}
