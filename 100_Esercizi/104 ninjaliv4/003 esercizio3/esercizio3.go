package main

import (
	"fmt"
)

func main() {

	numeri := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(numeri[:5])
	fmt.Println(numeri[5:])
	fmt.Println(numeri[2:7])
	fmt.Println(numeri[1:6])

}
