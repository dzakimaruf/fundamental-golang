package main

import "fmt"

func main() {
	//slice1 := []string{
	//	"Monkey",
	//	"D.",
	//	"Luffy",
	//}
	slice1 := make([]string, 3)
	slice1[0] = "Monkey"
	slice1[1] = "D."
	slice1[2] = "Luffy"

	slice2 := append(slice1, "budi", "nugraha")

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "Joko"

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]string, 2)

	copy(slice3, slice1)

	fmt.Println(slice1)
	fmt.Println(slice3)

	slice1[0] = "Budi"

	fmt.Println(slice1)
	fmt.Println(slice3)
}
