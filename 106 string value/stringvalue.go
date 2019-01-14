package main

import (
	"fmt"
)

func main() {

	s := "hi everybody"

	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//`= si inserisce con alt 0096 e si usa per spezzare a capo la stringa
	s1 := `hi
everybody`

	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	nbytes := []byte(s)
	nbytes1 := []byte(s1)
	// pur avendo uno spazio i caratteri non cambiano (qua sono stati convertiti in byte vedere tabella UTF 8)
	fmt.Println(nbytes)
	fmt.Println(nbytes1)
	fmt.Printf("%T\n", nbytes)
	fmt.Printf("%T\n", nbytes1)
}
