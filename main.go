package main

import "fmt"

func main() {
	//manifest type
	var firstname string = "Dzaki"

	//type inferences struct
	lastname := "Abdurrahman"
	lastname = "Maruf"

	//multi variable declaration
	var in, out, on string = "didalam", "diluar", "diatas"

	//pointer new
	name := new(string)

	//operator aritmatik
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Println("Hello", firstname, lastname, in, out, on, *name)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	//seleksi kondisi
	var point = 5

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//variabel temporary
	var point1 = 8840.0

	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	//switch case
	var point2 = 4

	switch point2 {
	case 10, 9, 8:
		fmt.Println("perfect")
	case 7, 6, 5:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}
	//switch case with if-else style
	var point3 = 3

	switch {
	case point3 == 8:
		fmt.Println("perfect")
	case (point3 < 8) && (point3 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//perulangan for
	for i := 1; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//perulangan dengan variabel
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	//perulangan tanpa argument

	var e = 0

	for {
		fmt.Println("Angka", e)

		e++
		if e == 5 {
			break
		}
	}
	//array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names)

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t:", len(fruits))
	fmt.Println("Isi semua element \t:", fruits)

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)
	
	//map
	
var chicken map[string]int
chicken = map[string]int{}

chicken["januari"] = 50
chicken["februari"] = 40

fmt.Println("januari", chicken["januari"]) // januari 50
fmt.Println("mei",     chicken["mei"])     // mei 0

var numberss1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numberss2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

fmt.Println("numbers1", numberss1)
fmt.Println("numbers2", numberss2)

var fruitss = [4]string{"apple", "grape", "banana", "melon"}

for i := 0; i < len(fruitss); i++ {
    fmt.Printf("elemen %d : %s\n", i, fruitss[i])
}

var frui = []string{"apple", "grape", "banana", "melon"}
fmt.Println(frui[0]) // "apple"
}

func main() {
    var names = []string{"John", "Wick"}
    printMessage("halo", names)
}

func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}