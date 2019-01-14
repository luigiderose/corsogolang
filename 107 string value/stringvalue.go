package main

import (
	"fmt"
)

func main() {

	s := "hi everybody" //non è mutabile non può essere cambiata

	fmt.Println(s)
	fmt.Printf("%T\n", s)
	//`= si inserisce con alt 0096 e si usa per spezzare a capo la stringa
	s1 := `hi
everybody`

	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	nbytes := []byte(s)   //array convertito in decimale della parola s
	nbytes1 := []byte(s1) //array convertito in decimale della parola s1
	// pur avendo uno spazio i caratteri non cambiano (qua sono stati convertiti in byte vedere tabella UTF 8)
	fmt.Println(nbytes)
	fmt.Println(nbytes1)
	fmt.Printf("%T\n", nbytes)
	fmt.Printf("%T\n", nbytes1)

	//*****************    visualizzazione esadecimale della stringa 1^ visualizzazione rappresentazione UTF 8  **********************//
	for i := 0; i < len(s); i++ {

		fmt.Printf("%#U", s[i])
	}

	//*****************    visualizzazione esadecimale della stringa 2^ visualizzazione rappresentazione pura esadecimale **********************//
	fmt.Println("")

	for i, v := range s {

		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
