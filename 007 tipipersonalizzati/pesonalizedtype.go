package main

import "fmt"

var a int //abbiamo dichiarato la variabie a come tipo intero

type gianni int //qua creiamo un tipo personalizzato che si chima gianni ed è di tipo intero

var b gianni //dichiariamo che la variabile b è di tipo gianni

func main() {

	a = 42
	b = 43

	fmt.Println(a)
	fmt.Printf("%T\n", a) //quando stampiamo il tipo di a ci uscirà int
	fmt.Println(b)
	fmt.Printf("%T\n", b) //quando stampiamo il tipo di b ci uscità main.gianni che è dove è posizionato la dichiarazione del tipo gianni e che tipo è

	/*successivamente visto che il linguaggio di go è un linguaggio statico non potremmo assegnare ad a il valore di b in quanto sono di due tipologie diverse
	quindi

	a=b
	fmt.Println(a) // da errore

	quindi cosa faremo dovremmo procedere alla conversione della tipologia
	*/

	a = int(b) //negli altri linguaggi è chiamato cast in golang è chiamato conversion
	fmt.Println(a)

}
