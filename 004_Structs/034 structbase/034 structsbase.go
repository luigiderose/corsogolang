package main

import (
	"fmt"
)

/*dichiaro una struttura oggetto che mette assieme oggetti di tipologia diversa*/
type person struct {
	name    string
	surname string
	age     int
}

func main() {
	x := person{"azzeddo", "mobili", 31} //popolo con dei dati la struttura
	fmt.Println("structs test")          //titolo
	fmt.Println(x)                       //stampo il record
	fmt.Println(x.name)                  //stampo il nome del record
	fmt.Println(x.surname)
	fmt.Println(x.age)
	fmt.Println(x.age, x.name)
	sx := &x
	fmt.Println(sx.age + 1) //è possibile effettuare operazioni
}
