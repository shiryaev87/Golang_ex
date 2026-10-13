package main

import "fmt"

var c, python, java bool

func main() {
	var i, j int = 1, 2
	k := 3
	c1, p1, j1 := true, false, "no!"
	fmt.Println(split(17))

	fmt.Println(i, c, python, java)

	fmt.Println(i, j, k, c1, p1, j1)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
