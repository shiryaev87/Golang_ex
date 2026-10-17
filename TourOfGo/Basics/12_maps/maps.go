package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var mp = map[string]Vertex{
	"Bell Labs": {
		40.68, -74.49,
	},
	"Google": {
		37.42, -122.08,
	},
}

func main() {
	fmt.Printf("hello maps\n")

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.6843, -74.399967,
	}

	fmt.Println(m["Bell Labs"])

	fmt.Println("--------MAps Literals-------")

	fmt.Println(mp)

	fmt.Println("--------Mutating maps--------")

	m1 := make(map[string]int)

	m1["Answer"] = 42
	fmt.Println("The value:", m1["Answer"])

	m1["Answer"] = 48
	fmt.Println("The value:", m1["Answer"])

	delete(m1, "Answer")
	fmt.Println("The value:", m1["Answer"])

	v, ok := m1["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println("The value:", m1)
}
