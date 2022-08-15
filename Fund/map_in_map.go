package main

import "fmt"

func main() {
	//mahasiswa := map[string]string{
	//	"1001" : "Monkey",
	//	"1002" : "D.",
	//	"1003" : "Luffy",
	//}

	mahasiswa := map[string]map[string]string{
		"1001": map[string]string{
			"name":    "Monkey",
			"address": "Indonesia",
			"gender":  "Male",
		},
		"1002": map[string]string{
			"name":    "D.",
			"address": "Indonesia",
			"gender":  "Male",
		},
		"1003": map[string]string{
			"name":    "Luffy",
			"address": "Indonesia",
			"gender":  "Male",
		},
	}

	delete(mahasiswa, "1003")

	fmt.Println(mahasiswa)
}
