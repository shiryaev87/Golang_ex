package main

import (
	"fmt"

	evenorodd "golang_task/codewars/8kyu/even_or_odd"
)

func main() {
	// for _, n := range []int{0, 1, 2, 3, -4} {
	// fmt.Printf("%d -> %s\n", n, evenorodd.EvenOrOdd(n))
	// }

	fmt.Printf("%s", evenorodd.EvenOrOdd(7))
}
