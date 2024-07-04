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

	// Type assertion
	var any interface{}
	any = "hello"
	str := any.(string)
	fmt.Println(str)

	// interface to map
	var any2 interface{}
	any2 = make(map[string]string)
	any2.(map[string]string)["name"] = "david"

	fmt.Printf("%T\n", any2)
	fmt.Println(any2)
	// below is invalid
	// fmt.Println(any2["name"])

	// interface to map with intermediate variable
	var any3 interface{}
	any3 = make(map[string]string)
	map3 := any3.(map[string]string)
	map3["name"] = "mike"
	// both print the same thing
	fmt.Println(any3)
	fmt.Println(map3)

	// interface to map of interface
	var any4 interface{}
	any4 = make(map[interface{}]interface{})
	map4 := any4.(map[interface{}]interface{})
	map4[3] = "name"
	fmt.Println(map4)
	fmt.Println(map4[3].(string))

}
