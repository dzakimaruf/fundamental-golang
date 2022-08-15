package main

import "fmt"

func main() {
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Monkey"
	mahasiswa["1002"] = "D."
	mahasiswa["1003"] = "Luffy"

	fmt.Println(mahasiswa)

	fmt.Println(mahasiswa["1001"])
	fmt.Println(mahasiswa["1002"])
	fmt.Println(mahasiswa["1003"])

	for nim, name := range mahasiswa {
		fmt.Println("Nim", nim, "Bernama", name)
	}
}
