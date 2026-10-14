package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices")

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println("выводим слайсы  a b", a, b)

	b[0] = "XXX" // меняем занчение у слайса b 0

	fmt.Println(a, b)

	fmt.Println(names)
}
