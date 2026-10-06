package main

import "fmt"

func main() {
	x := make([]int, 0, 2)
	x = appendInt(x, 1)
	fmt.Printf("x = %v, len = %d, cap = %d\n", x, len(x), cap(x))
	x = appendInt(x, 2)
	fmt.Printf("x = %v, len = %d, cap = %d\n", x, len(x), cap(x))
	x = appendInt(x, 3)
	fmt.Printf("x = %v, len = %d, cap = %d\n", x, len(x), cap(x))
	x = appendInt(x, 3)
	fmt.Printf("x = %v, len = %d, cap = %d\n", x, len(x), cap(x))
	x = appendInt(x, 3)
	fmt.Printf("x = %v, len = %d, cap = %d\n", x, len(x), cap(x))
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	// анализируем расширять срез или нет`	`
	if zlen <= cap(x) {
		// есть место куда добавлять элемент
		z = x[:zlen]
	} else {
		// места для роста нет. выделяем новый массив
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
