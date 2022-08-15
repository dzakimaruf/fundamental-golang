package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(1, 2, 3)
	fmt.Println(total)	
}

// func main() {
//     var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
//     var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
//     fmt.Println(msg)
// }

// func calculate(numbers ...int) float64 {
//     var total int = 0
//     for _, number := range numbers {
//         total += number
//     }

//     var avg = float64(total) / float64(len(numbers))
//     return avg
// }