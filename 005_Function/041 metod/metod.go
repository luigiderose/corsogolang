package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type stato struct {
	person
	statos string
}

func (p stato) stato() {
	fmt.Println("sono", p.first, p.last, "sono stato", p.statos)
}

func main() {

	pp1 := stato{
		person: person{
			first: "Luigi",
			last:  "De Rose",
		},
		statos: "PROMOSSO",
	}

	pp2 := stato{
		person: person{
			first: "Vanessa",
			last:  "Frau",
		},
		statos: "PROMOSSO",
	}

	pp3 := stato{
		person: person{
			first: "Lorenzo",
			last:  "De ROse",
		},
		statos: "BOCCIATO",
	}
	fmt.Println(pp1)
	pp1.stato()
	pp2.stato()
	pp3.stato()
}
