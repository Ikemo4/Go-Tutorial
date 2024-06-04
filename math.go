package main

import (
	"fmt"
	"math/rand"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(split(17))
}