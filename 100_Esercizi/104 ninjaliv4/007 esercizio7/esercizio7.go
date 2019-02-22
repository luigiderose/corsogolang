package main

import (
	"fmt"
)

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	name := [][]string{x, y}
	fmt.Println(name)

	for i, v := range name {
		fmt.Println("Record: ", i)
		for j, z := range v {
			fmt.Printf("\t data:  \t %v %v\n", j, z)
		}
	}

}
