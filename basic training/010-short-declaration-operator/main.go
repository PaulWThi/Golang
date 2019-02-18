package main

import (
	"fmt"
)

var y = 10
var z int

func main() {
	x := 20
	fmt.Println(x)
	fmt.Println(y)
	foo()
	z = 30
	fmt.Println(z)
}
func foo() {
	fmt.Println("again: ", y)
}
