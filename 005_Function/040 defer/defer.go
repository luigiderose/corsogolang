package main

import (
	"fmt"
)

func main() {
	defer foo() //serve a passare per ultimo questa funzione
	bar()
	pippo()
}

func foo() {

	fmt.Println("foo")
}

func bar() {

	fmt.Println("bar")
}

func pippo() {
	fmt.Println("pippo")
}
