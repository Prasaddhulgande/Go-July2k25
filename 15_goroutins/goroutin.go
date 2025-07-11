package main

import (
	"fmt"
	"sync"
)

func task(id int, W *sync.WaitGroup) {
	defer W.Done()
	fmt.Println("Doing Task", id)
}

// Use of go routins, Multipule time run programs, Multithreading, Paralal threading on seperate thread.
// Useing "go" key to use goroutins. Powerfull feature of golang.
func main() {

	var wg sync.WaitGroup // wait group

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg) //OR

		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}

	// time.Sleep(time.Second * 1)
	wg.Wait()
}
