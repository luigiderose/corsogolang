package main

import (
	"fmt"
)

/*dichiarando come variabili o tipi esterni li posso utilizzare nella main o in funzioni richiamate main*/

type person struct {
	FirstName string
	LastName  string
	IceCream  []string
}

/*Variabili Persone*/
var p1 = person{
	FirstName: "Luigi",
	LastName:  "De Rose",
	IceCream:  []string{"chocolate", "cream", "strawberris"},
}

var p2 = person{
	FirstName: "Azzeddo",
	LastName:  "Haida",
	IceCream:  []string{"strawberris", "grapefruit", "banana"},
}

/*Mappa Utenti*/
var utenti = map[int]person{
	0: p1,
	1: p2,
}

/*Funzioni*/
func test() {

	for _, z := range utenti {
		fmt.Printf("%v %v\n", z.FirstName, z.LastName)
		for i, val := range z.IceCream {
			fmt.Printf("\t\t%v\t%v\n", i, val)
		}
	}
}

/*main*/
func main() {

	test()

}
