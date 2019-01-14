package main

import "fmt"

const (
	_  = iota             //dichiaro una costante vuota che non utilizzo valore iota
	kb = 1 << (iota * 10) // ora prendo il bite e lo shifto di 10 volte verso sinistra iota per 10 quindi 1000000000
	mb = 1 << (iota * 10) // incremento di dieci volte iota quindi altre 10 posizioni 100000000000000000000
	gb = 1 << (iota * 10) // e così via
	tb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}
