package main

import (
	"fmt"
)

type utente struct { //crea una struttura
	id      int
	nome    string
	cognome string
	age     int
}

type promosso struct { //crea una seconda struttura
	utente
	stato bool
}

func main() {

	x := promosso{ //popola la struttura nella struttura
		utente: utente{1, "Luigi", "De Rose", 32},
		stato:  true,
	}

	y := promosso{ //popola la struttura nella struttura
		utente: utente{2, "Vanessa", "Frau", 30},
		stato:  false,
	}

	fmt.Printf("utente: %v\t\tpromosso: %v\n", x.cognome, x.stato) //stampa di dati richiesti
	fmt.Printf("utente: %v\t\tpromosso: %v\n", y.cognome, y.stato)
	fmt.Printf("%T", x)

}
