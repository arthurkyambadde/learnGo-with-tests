package summation

// func Sum(numbers [5]int) int {

// 	var total int = 0

// 	for i := 0; i < len(numbers); i++ {
// 		total = total + numbers[i]
// 	}

// 	return total
// }

func Sum(numbers [5]int) int {

	total := 0
	for _, number := range numbers {
		total += number

	}
	return total
}
