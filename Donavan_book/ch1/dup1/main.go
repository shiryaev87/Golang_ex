// dup1 выводит повторяющиеся строки, которые появляются в стандартном вводе, вместе с количеством их повторений.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// примечание. игнорируем возможные ошибки ввода
	// ошибки из input.Err() можно обрабатывать здесь, если необходимо
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
