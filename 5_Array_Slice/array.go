package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num [4]int

	fmt.Println(len(num))

	num[1] = 1
	num[2] = 2

	fmt.Println(num)

	// In integer case, if we are not provided any values then print Zero
	// but in boolean type array, print false all values

	Nums := [3]int{11, 22, 33}

	fmt.Println(Nums)

	// 2D arrays

	nums := [4][3]int{{2, 4}, {1, 3}}

	fmt.Println(nums)

	// Fixed size, that is predictable
	// memory optimization
	// Costant time access, Quick, Fast access

	//slices()

	CopySlice()
}

func slices() {

	// more dynamic,  most used construct in go, useful methods

	var nums []int
	//by default nil values print
	fmt.Println(nums)

	nnums := []int{1, 2, 3, 4, 5}
	fmt.Println(nnums)

	nnums = append(nnums, 6)
	nnums = append(nnums, 7)

	fmt.Println(nnums)
}

func CopySlice() {

	var nums = make([]int, 0, 5)

	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	var nums2 = make([]int, len(nums))

	copy(nums2, nums)

	fmt.Println(nums, nums2)

	// Compare two slices

	fmt.Println(reflect.DeepEqual(nums, nums2))

	nums2 = append(nums2, 7)

	fmt.Println(reflect.DeepEqual(nums, nums2))

}
