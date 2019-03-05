package main

import (
	"fmt"
)

func main() {

	foo()
	bar("James")
	s1 := somma(5, 6)
	fmt.Println(s1)
	s2, s3 := moltiplicadividi(4, 2)
	fmt.Println(s2, s3)
}

/*

Come si scrive una funzione?

func (r receiver) identifier(parameters) (return(s)){...}

es sotto la func foo

tutti i valori vengono passati per valore by value

es 2 func bar

puo ritornare dei valori

es 3 somma di 2 numeri
*/

func foo() {
	fmt.Println("pippo")
}

func bar(s string) {
	fmt.Println("hello, ", s)
}

func somma(a int, b int) int {
	y := a + b
	return y
}

func moltiplicadividi(a int, b int) (int, int) {
	y := a * b
	z := a / b
	return y, z
}
