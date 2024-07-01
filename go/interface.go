package main

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) Bark() {
	fmt.Println("W")
}

func (c Cat) Bark() {
	fmt.Println("C")
}

func main() {
	var d = Dog{"dog"}
	var c = Cat{"cat"}
	var in_d = Animal(d)
	var in_c = Animal(c)
	in_d.Bark()
	in_c.Bark()
}
