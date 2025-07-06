package main

import "fmt"

func main() {
	whileloop()
	clasicFOR()

	

}

// While loop

func whileloop() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}

func clasicFOR() {

	for i := 1; i <= 10; i++ {

		//break

		// Skipe 2,4,6 & continue other numbers
		if i == 2 || i == 4 || i == 6 {
			continue
		}
		fmt.Println(i)
	}
}

//range loop



	

