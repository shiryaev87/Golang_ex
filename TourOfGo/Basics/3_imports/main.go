package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %e problems. \n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Printf("сумма %d и %d  равно %d \n", 10, 20, add(10, 20))

	fmt.Println(swap("Hello", "World"))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
