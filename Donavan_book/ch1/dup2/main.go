// dup1 выводит повторяющиеся строки, которые появляются в стандартном вводе, вместе с количеством их повторений.
// программа читает ввод или список именованных файлов
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // создаем мапу для хранения количества потворений ключ значение
	//input := bufio.NewScanner(os.Stdin) // создаем переменную input,  которая будет сканировать стандартный ввод
	files := os.Args[1:] // создаем переменную files, которая будет хранить список именованных файлов, переданных в аргументах командной строки
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	} // если список файлов пустой

	// примечание. игнорируем возможные ошибки ввода
	// ошибки из input.Err() можно обрабатывать здесь, если необходимо
	for line, n := range counts { // запускаем цикл, который будет перебирать строки и количество их повторений
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// примечание . игнорируем потенциальные ошибки ввода из input.err(), если необходимо, их можно обрабатывать здесь
}
