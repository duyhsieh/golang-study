package main

import "fmt"

var gVar int = 0
func main() {
	gVar+=1
	fmt.Println("main:global var=", gVar)
	foo()
}

func foo() {
	gVar+=1
	fmt.Println("foo:global var=", gVar)
}
