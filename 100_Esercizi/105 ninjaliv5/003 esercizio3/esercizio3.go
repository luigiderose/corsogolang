package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type luxury struct {
	vehicle
	luxury bool
}

var v1 = truck{
	vehicle: vehicle{
		doors: "two",
		color: "red",
	},
	fourWheel: true,
}

var v2 = luxury{
	vehicle: vehicle{
		doors: "four",
		color: "blue",
	},
	luxury: true,
}

func main() {
	luigiFunction()

}

func luigiFunction() {
	fmt.Println(v1, v2)
	fmt.Println(v1.doors)
	fmt.Println(v2.doors)

}
