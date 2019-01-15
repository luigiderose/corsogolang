package main

import (
	"fmt"
)

func main() {

	//dichiaro una stringa
	s := "H"

	//la stampo
	fmt.Println(s)

	//la converto in byte
	bs := []byte(s)

	fmt.Println(bs)

	// estrapolo il numero nell'array 0
	n := bs[0]

	//stampo il tipo di dato
	fmt.Printf("%T, tipo di dato \n", n)
	// stampo il numero con il %v valore di origine
	fmt.Printf("%v, numero\n", n)
	// stampo il numero in base decimale (%d)
	fmt.Printf("%d, numero\n", n)
	//stampo il numero in base binaria (%b)
	fmt.Printf("%b, numero binario\n", n)
	//stampo il numero in base esadecimale (%#X)
	fmt.Printf("%#X, esadecimale\n", n)
	//stampo il numero con codifica UTF 8
	fmt.Printf("%#U, UTF-8 \n", n)
}
