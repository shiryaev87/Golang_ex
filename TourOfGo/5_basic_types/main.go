package main

import (
	"fmt"
	"math/cmplx"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

var (
	ToBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// var i, x, y int
	// var f float64
	// var b bool
	// var s string = "aasdad"

	// fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value %v\n", maxInt, maxInt)
	// fmt.Printf("Type: %T Value %v\n", z, z)

	// fmt.Printf("%v %v %v %q", i, f, b, s)

	// x, y = 3, 4
	// f = float64(math.Sqrt(float64(x*x + y*y)))

	// z := uint(f)
	// fmt.Printf("%v %v %v", x, y, z)

	fmt.Println(needInt(Small))

	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(z float64) float64 { return z * 0.1 }
