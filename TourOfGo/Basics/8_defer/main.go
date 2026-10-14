package main

import "fmt"

func main() {
	defer fmt.Println("World")
	x := 1
	defer fmt.Printf("запомнил перед defer %d ", x)
	x = 2
	fmt.Printf("просто %d ", x)

	fmt.Println("hello")

	// Stacking defers

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
