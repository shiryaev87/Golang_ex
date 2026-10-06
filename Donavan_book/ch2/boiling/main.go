// boiling выводит температуру кипения воды в градусах Цельсия и Фаренгейта
package main

import "fmt"

const boilingF = 212.0

func main() {
	f := boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("Температура кипения воды = %gF, илл %gC\n", f, c)
	// вывод
	// Температура кипения = 212F или 100 C
}
