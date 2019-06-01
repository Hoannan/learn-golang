package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	level int
}

func (p *person) getStudent() {
	fmt.Println("student info:", *p)
}

func main() {
	khn := student{
		person{"khn", 25},
		109,
	}

	khn.person.getStudent()
}
