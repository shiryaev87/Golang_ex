package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v11 = Vertex{1, 2}
	v12 = Vertex{X: 1}
	v13 = Vertex{}
	p11 = &Vertex{1, 2}
)

func main() {
	fmt.Println("Hello")
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 30
	fmt.Println(v.X)

	v2 := Vertex{1, 2}
	p := &v2
	p.X = 1e9
	fmt.Println(v2)

	fmt.Println(v11, p11, v12, v13)
}

// если p — указатель на структуру, к её полям можно обращаться почти так же, как к полям самой структуры.

// Указатель p указывает на v. Чтобы добраться до поля X, можно явно разыменовать указатель:

// Но Go позволяет записать короче:

// В этом случае Go автоматически разыменовывает p. Поэтому p.X и (*p).X — одно и то же.

// Изменение поля через указатель изменит исходную структуру:
