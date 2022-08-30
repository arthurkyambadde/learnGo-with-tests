package summation

func Sum(numbers []int) int {

	total := 0
	for _, number := range numbers {
		total += number

	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
