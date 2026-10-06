package main

// Echo2 выводит аргументы командной строки, используя цикл for с функцией range

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
