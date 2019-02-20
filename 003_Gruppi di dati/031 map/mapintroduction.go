package main

import (
	"fmt"
)

func main() {
	/* maps è una tipologia di chiave che permette la localizzazione di dati in maniera efficente e veloce

	   si utilizza per convenzione la lettera m per la variabile m:= map[tipologia per esempio string]tipologiaarray{}

	*/
	m := map[string]int{ //per popolare la variabile map utilizzare la struttura "key":value,   la virgola separa i valori

		"Luigi":  32,
		"Vanesa": 30,
		"Luna":   4,
	}

	fmt.Println(m)

	/*se ora richiamo una key mi verrà restituito il valore corrispondente*/

	fmt.Println(m["Luna"])

	/* ma se richiamo un key che non c'è allora il mi verrà stampato il valore zero*/
	fmt.Println(m["Peppe"])

	/*come si fa per verificare pertanto se un oggetto è presente con l'identificatore di variabile "ok"*/

	v, ok := m["Peppe"]
	fmt.Println(v)  //questo mi stamperà il valore
	fmt.Println(ok) // questo mi stamperà falso

	/*quindi ecco in funzione un if*/

	if v, ok := m["Luna"]; ok {
		fmt.Println("l'età di Luna è: ", v)
	}

	z := map[int]string{ //ho testato lo stesso valore non è ammesso nella chiave (pertanto diventa univoca)
		32: "luigi",
		30: "vanessa",
		31: "lurenzo",
		4:  "luna",
	}

	fmt.Println(z)
}
