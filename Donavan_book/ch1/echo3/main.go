package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Echo3 выводит аргументы командной строки, используя функцию strings.Join
func main() {
	// ar s, sep string
	start := time.Now()
	fmt.Printf(strings.Join(os.Args[0:], "\n"))

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %d ns\n", elapsed)
}
