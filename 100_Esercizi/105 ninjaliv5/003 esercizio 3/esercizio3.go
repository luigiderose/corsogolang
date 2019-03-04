package main

import (
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	IceCream  []string
}

func main() {

	p1 := person{
		FirstName: "Luigi",
		LastName:  "De Rose",
		IceCream:  []string{"chocolate", "cream", "strawberris"},
	}

	p2 := person{
		FirstName: "Azzeddo",
		LastName:  "Haida",
		IceCream:  []string{"strawberris", "grapefruit", "banana"},
	}

	fmt.Printf("%v %v:\n", p1.LastName, p1.FirstName)
	for i := range p1.IceCream {
		fmt.Printf("\t\t%v\n", p1.IceCream[i])
	}

	fmt.Printf("%v %v:\n", p2.LastName, p2.FirstName)
	for v := range p2.IceCream {
		fmt.Printf("\t\t%v\n", p2.IceCream[v])
	}
}
