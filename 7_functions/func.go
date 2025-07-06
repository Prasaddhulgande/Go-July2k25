package main

import "fmt"

func main() {
	result := add(5, 8)
	fmt.Println(result)

	lag1, lag2, _ := getLanguages()
	fmt.Println(lag1, lag2)

	fn := proccessit()
	fn(6)
}

func add(a, b int) int {

	return a + b

}

func getLanguages() (string, string, string) {

	return "golang", "Java", "C#"
}

// Pass fun as a parameter

func proccessit() func(a int) int {
	return func(a int) int {
		return 4

	}
}
