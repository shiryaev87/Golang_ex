package main

import "fmt"

func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"a", "", "b", "", "", "c"}
	fmt.Println("До вызова", data)
	result := noempty(data)
	fmt.Println("Результат	", result)
	fmt.Println("После вызова", data)
}
