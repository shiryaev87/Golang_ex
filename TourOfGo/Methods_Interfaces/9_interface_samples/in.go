package main

import "fmt"

// 1. Создаем интерфейс (нашу вакансию)
type Courier interface {
	Deliver(address string)
}

// 2. Создаем первого кандидата (Студента)
type Student struct {
	Name string
}

// Студент умеет доставлять пиццу
func (s Student) Deliver(address string) {
	fmt.Printf("%s привез пиццу на велосипеде на адрес: %s\n", s.Name, address)
}

// 3. Создаем второго кандидата (Дрон)
type Drone struct {
	Model string
}

// Дрон тоже умеет доставлять пиццу
func (d Drone) Deliver(address string) {
	fmt.Printf("Дрон %s прилетел по воздуху на адрес: %s\n", d.Model, address)
}

// 4. Универсальная функция пиццерии.
// Она принимает ЛЮБОГО Courier. Ей неважно, человек это или робот.
func SendOrder(c Courier, address string) {
	c.Deliver(address)
}

func main() {
	ivan := Student{Name: "Иван"}
	dji := Drone{Model: "Mavic 3"}

	// Мы можем передать в функцию и Ивана, и Дрон,
	// потому что они оба неявно подходят под интерфейс Courier!
	SendOrder(ivan, "ул. Ленина, 10")
	SendOrder(dji, "ул. Пушкина, 25")
}
