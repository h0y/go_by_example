package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	s := person{name: "ellery", age: 20}

	fmt.Println(s)
	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp)
	fmt.Println((*sp).name)
	fmt.Println(sp.name)
}