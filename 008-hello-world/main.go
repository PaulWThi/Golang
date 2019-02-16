package main

import "fmt"

func main() {
	fmt.Println("hello everyone! lets git git")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	fmt.Println("I'm in foo baby!")
}
func bar() {
	n, _ := fmt.Println("and then we exited")
	fmt.Println(n)
}
