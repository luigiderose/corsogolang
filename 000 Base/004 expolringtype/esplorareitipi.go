package main

import (
	"fmt"
)

// primitive data type
var y = 42

// static programming language
var z = "pippo pedro zuzzo"                //dichiarata come stringa
var a = `pedro dice "gianni sono felice "` // con l'accento grave vogliamo dire al programma che le virgolette fanno parte del testo e non lo delimitano

// altgr 0096 = `` mentre per ~ altgr 0126 (su questa tastiera alt sinistro)

// composite data type
// esmpio = array, slice, struct
func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y) //mostra il tipo della variabile selezionata (int)
	fmt.Println(z)
	fmt.Printf("%T\n", z) //mostra il tipo della variabile selezionata (string)
	fmt.Println(a)
	// se provassi ad assegnare a z = 43 quindi un valore int avrei un errore}
}
