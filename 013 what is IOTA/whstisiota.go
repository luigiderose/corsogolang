package main

import (
	"fmt"
)

//iota è una predichiarazione che auto incrementa di 1 i valori

const (
	a = iota
	b = iota
	c = iota
)

//anche in questo modo iota è una predichiarazione che auto incrementa di 1 i valori ogni colta che dichiaro una suddivisione delle costanti riparte da 0
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
