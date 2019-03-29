package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type agenteSegreto struct {
	person
	agente bool
}

var persona1 = agenteSegreto{
	person: person{
		first: "Luigi",
		last:  "De Rose",
	},
	agente: true,
}

var persona3 = agenteSegreto{
	person: person{
		first: "Vanessa",
		last:  "Frau",
	},
	agente: true,
}

var persona2 = person{
	first: "Michelle",
	last:  "Frau",
}

func (p person) rivelaSegreti() {
	fmt.Println(p.first, p.last, "ho detto un segreto")
}

type esseriUmani interface {
	rivelaSegreti()
}

func bar(e esseriUmani) {
	//assering
	switch e.(type) {
	case person:
		{
			fmt.Println("io non sono un agente segreto", e.(person).first)
		}
	case agenteSegreto:
		{
			fmt.Println("io SONO un agente segreto", e.(agenteSegreto).first)
		}
	}
	fmt.Println("sono passato dal bar", e)
}

func scrivipersone() {
	fmt.Println("ORA SCRIVO I NOMI DELLE PERSONE")
	fmt.Println(persona1.first, persona2.first, persona3.first)
	fmt.Printf("\n")
}

func scrivisegreti() {
	fmt.Println("ORA SCRIVO CHI DICE I SEGRETI")
	persona2.rivelaSegreti()
	persona3.rivelaSegreti()
	fmt.Printf("\n")
}

func scrivibar() {
	fmt.Println("ORA SCRIVO CHI PASSA DAL BAR")
	bar(persona1)
	bar(persona2)
	bar(persona3)
	fmt.Printf("\n")
}

func conversion() {
	fmt.Println("QUESTA E' LA FUNZIONE DI CONVERSIONE")
	type numero int

	y := int(57)
	z := numero(y)
	fmt.Println("Scrivo un numero")
	fmt.Println(y)
	fmt.Println("converto il tipo")
	fmt.Println(z)
	fmt.Println("mostro di che tipo è il primo")
	fmt.Printf("%T\n", y)
	fmt.Println("mostro di che tipo è diventato")
	fmt.Printf("%T\n\n", z)

}

func main() {
	scrivipersone()
	scrivisegreti()
	scrivibar()
	conversion()
	fmt.Println("FINE")

}
