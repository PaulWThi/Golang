package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// %v prints value
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Printf(s)
}
