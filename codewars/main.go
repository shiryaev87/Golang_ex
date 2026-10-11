// Finish the solution so that it sorts the passed in array of numbers. If the function passes in an empty array or null/nil value then it should return an empty array.

// For example:

// solution([1, 2, 3, 10, 5]) --> should return [1, 2, 3, 5, 10]
// solution(null)             --> should return []
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortNumbers([]int{1, 2, 3, 10, 5}))
}

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	sort.Ints(numbers)
	return numbers
}
