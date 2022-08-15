package main

import "fmt"

type Address struct {
	City, Country string
}

func main() {
	

var address1 Address = Address{"Bogor", "Indonesia"}

var address2 := &address1

address2.City = "Cianjur"

fmt.Println(address1)
fmt.Println(address2)
}