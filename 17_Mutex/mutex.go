package main

import (
	"fmt"
	"sync"
)

type post struct {
	Views int

	// Define mutex, use of mutex is to lock that perticular goroutins
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {

	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.Views += 1

}

func main() {
	var wg sync.WaitGroup
	myPost := post{Views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg) // at a time running all & print all numbers with the help of goroutins
		// Calling all veiw at a time.
	}

	// myPost.inc()
	fmt.Println(myPost.Views)
}
