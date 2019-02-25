package main

import (
	"fmt"
)

func main() {
	x := struct { //struttura anonima perchè dichiarata all'interno senza creare un type
		nome    string
		cognome string
		età     int
	}{
		nome:    "luigi",
		cognome: "derose",
		età:     32}

	fmt.Println(x)

}
