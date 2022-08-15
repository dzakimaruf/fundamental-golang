package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var dzaki Customer
	dzaki.Name = "dzaki"
	dzaki.Address = "Cijahe"
	dzaki.Age = 27
}