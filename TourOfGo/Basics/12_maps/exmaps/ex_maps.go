package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
	//"golang.org/x/tour/wc"
)

func main() {
	fmt.Println("--------Maps ex------")
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	arr_s := strings.Fields(s)
	fmt.Println(arr_s)
	for _, word := range strings.Fields(s) {
		counts[word]++
	}
	return counts
}
