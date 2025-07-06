package main

import "fmt"

func changeNum(num *int) {

	*num = 5

	fmt.Println("In ChnageNum", *num)
}

func main() {
	num := 1
	fmt.Println("In main", num)
	
	changeNum(&num)

	fmt.Println("After change in main", num)
}
