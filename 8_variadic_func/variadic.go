package main

import "fmt"

func main() {

	res := sum(1, 2, 3, 4, 5)

	fmt.Println(res)
	// fmt.Println(sum())

	num := []int{6, 7, 8, 9}

	ress := sum(num...)
	fmt.Println(ress)

}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {

		total = total + num
	}
	return total
}
