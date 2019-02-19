package main

import (
	"fmt"
)

const a int32 = 42
const b = 42

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T,%T", a, b)
}
