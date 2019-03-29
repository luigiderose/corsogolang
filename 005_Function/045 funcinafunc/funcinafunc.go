package main

import (
	"fmt"
)

func main() {

	fmt.Println(bar()())
	fmt.Printf("%T", bar())
}

func bar() func() string {
	return func() string {
		return "pippo"
	}
}
