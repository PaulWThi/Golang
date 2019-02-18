package main

import (
	"fmt"
)

var y string
var z int

func main() {
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Thi, Paul Thi"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T", z)
}
