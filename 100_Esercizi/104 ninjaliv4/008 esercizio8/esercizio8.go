package main

import "fmt"

func main() {
	x := map[string][3]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	for i, v := range x {
		fmt.Println("Last_firstname: ", i)
		for _, z := range v {
			fmt.Printf("\t \t \t \t %v\n", z)
		}

	}

}
