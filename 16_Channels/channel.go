package main

import (
	"fmt"
)

// func processNum(numChan chan int) {

// 	for num := range numChan { //infinite loop
// 		fmt.Println("Processing Number", num)
// 		time.Sleep(time.Second)
// 	}

// }

func sum(result chan int, num1, num2 int) {
	NumResult := num1 + num2
	result <- NumResult
}

// Without wait group run blocking run channels
// Goroutin synchronization
func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("Processing.....")
}
func main() {

	// chan1 := make(<-chan int)  // Received-only type channel, can't send data on this channels
	// chan2 := make(chan<- string)  // Send type channel, can't receive data from this channel.

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 1
	}()

	go func() {
		//Send data to the channels
		chan2 <- "Pong"
	}()

	for i := 1; i <= 2; i++ {

		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2", chan2Val)
		}
	}

	/*// Un-Buffer channels --> limited amount off we can send
	emailChan := make(chan string, 100)
	emailChan <- "1@gmail.com"
	emailChan <- "2@gmail.com"
	emailChan <- "3@gmail.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan) */

	// This is buffer channels --> Means at one time send one value & receive also same.
	/* done := make(chan bool)
	go task(done)
	<-done  */ //block

	/* result := make(chan int)

	go sum(result, 4, 5)

	res := <-result //blocking

	fmt.Println(res)  */

	/*	numChan := make(chan int)

		go processNum(numChan)

		// numChan <- 5
		for { // infinite for loop
			numChan <- rand.Intn(100)

		} */

	// time.Sleep(time.Second * 2)  // No need of this thread, Because of this loop is infinite

	// messageChan := make(chan string)

	// messageChan <- "Ping in to the channels"

	// msg := <-messageChan

	// fmt.Println(msg)
}
