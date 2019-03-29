package main

import (
	"fmt"
)

func main() {
	x := sum(2, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("il totale è ", x)
}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
		fmt.Println("il numero sommatto adesso è ", sum)
	}
	return sum
}
