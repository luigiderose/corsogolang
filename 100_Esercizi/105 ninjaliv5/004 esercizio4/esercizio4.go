package main

import "fmt"

func main() {
	s := struct {
		nome    string
		cognome string
	}{
		nome:    "luigi",
		cognome: "de rose",
	}

	fmt.Println(s)
}
