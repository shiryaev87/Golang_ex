// Consider an array/list of sheep where some sheep may be missing from their place. We need a function that counts the number of sheep present in the array (true means present).

// For example,

// [true,  true,  true,  false,
//   true,  true,  true,  true ,
//   true,  false, true,  false,
//   true,  false, false, true ,
//   true,  true,  true,  true ,
//   false, false, true,  true]
// The correct answer would be 17.

// Hint: Don't forget to check for bad values like null/undefined

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountSheeps([]bool{true, false, true, true, true, false}))
}

func CountSheeps(numbers []bool) int {
	if len(numbers) == 0 {
		return 0
	}
	count := 0
	for _, n := range numbers {
		if n == true {
			count++
		}
	}
	return count
}
