package main

import (
	"fmt"
	"runtime"
)

var z int8 = -128
var x int
var y float32

func main() {

	x = 42
	y = 42.30145

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
