package integers

import "fmt"

func Add(x, y int) int {
	//Add function adds two integers and returns there total/sum
	return x + y
}

func integers() {
	Add(2, 2)

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
