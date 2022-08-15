package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	Monkey := Man{"Monkey"}
	Monkey.Married()

	fmt.Println(Monkey.Name)
}
