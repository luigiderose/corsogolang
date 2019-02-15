package main

import (
	"fmt"
)

func main() {
	//to make an array
	var x [5]int   //we need to specify the dimension in this case is  5 position 0,1,2,3,4.
	fmt.Println(x) //print the contents of the array in 5 position 0.0.0.0.0

	x[2] = 2       // i can assing a value in the position of the array
	fmt.Println(x) // the result is [0.0.2.0.0]

}
