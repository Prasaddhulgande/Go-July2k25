package main

import "fmt"

// any/interface/comparable--> is a generic key to allow any type of date like interface{}
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// use generic with struct
type stack[T any] struct {
	elements []T
}

func main() {

	// nums := []int{1, 2, 3}
	names := []string{"Golang", "typeScript"}
	// func printSlice(items[]string)

	printSlice(names)

	myStack := stack[string]{
		elements: []string{"GO", "Java"},
	}

	fmt.Println(myStack)
}
