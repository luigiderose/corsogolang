package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang spec")

	/*CONSULTABILE SUL SITO GOLANG.ORG/REF/SPEC

	  ****** IDENTIFICATORI *********
	  come dichiarare una variabile

	  struttura

	  identifier = letter { letter | unicode_digit }

	  a
	  _ass
	  pippofranchino

	  se vogliamo specificare la tipoligia di variabile essa va dichiarata (vedere #Predeclared_identifiers)

	  ****** KEY WORD **********
	  ci sono alcuen parole chiave che non possono essere utilizzate come varibili in quanto sono utilizzate dal linguaggio GO

	  ***** operatori ***********
	  := breve dichiara e assegna
	*/

	x := 150 //dichiara x e assegna il valore 150
	fmt.Println(x)
	x = 1 //x è già dichiarato ma gli assegna il valore 1
	fmt.Println(x)

	y := 159 + 1 //y in questo caso è dichiarata come una variabile creata da un espressione
	fmt.Println(y)
}
