package main

import (
	"fmt"
)

const (
	a = (iota + 2019)
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
