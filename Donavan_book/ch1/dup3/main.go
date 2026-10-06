// dup3
package main

import (
	//	"bufio"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int) // создаем мапу для хранения количества потворений ключ значение
	//input := bufio.NewScanner(os.Stdin) // создаем переменную input,  которая будет сканировать стандартный ввод
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

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
