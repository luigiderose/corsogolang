package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
		"zio_peppe":       []string{"delfini", "unicornim", "Scoiattoli"}, //aggiunto internamente
	}

	x["lorenzo"] = []string{"Copioni", "triceratopi", "suricate"} //aggiunto esternamente

	for i, v := range x {
		fmt.Println("Last_firstname:", i)
		for _, z := range v {
			fmt.Printf("\t \t \t \t %v\n", z)
		}

	}

}
