package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	sum := 0
	for _, num := range nums {

		sum = sum + num
		// fmt.Println(num)

	}

	fmt.Println(sum)

	//Maps

	m := map[string]int{"ID": 11, "phone": 121211, "pin": 214309}

	for K, v := range m {

		fmt.Println(K, v)
	}

	for i, c := range "golang" {

		fmt.Println(i, c)
// Printing index and unicodes 
		fmt.Println(i, c)

	}

}
