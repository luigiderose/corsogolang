package main

import (
	"fmt"
)

var y = 42

func main() {

	fmt.Println(y) //stampa la linea

	//in ogni caso guardare la guida codice su golang.org fmt

	fmt.Printf("%T\n", y)  // prima della parentesi stampa formattato la parte prima della virgola %T vuol dire tipologia \n vuol dire a capo dopo la vigola di che cosa
	fmt.Printf("%b\n", y)  // %b numero binario
	fmt.Printf("%x\n", y)  // %x numero esadecimale
	fmt.Printf("%X\n", y)  // %X esadecimale maiuscolo
	fmt.Printf("%#x\n", y) //%#x alternato con esadecimale aggiunge 0x davanti

	fmt.Printf("%d\t%b\t%x", y, y, y) // stampa la y in tre modi e sposta i risultati di un tab \t
}
