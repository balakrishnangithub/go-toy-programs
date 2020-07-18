// credit: https://yourbasic.org/golang/data-races-explained/
// https://yourbasic.org/golang/gotcha-data-race-closure/
// A data race occurs when two goroutines access the same variable concurrently
// and at least one of the accesses is a write.
// $ go run --race .
package main

import (
	"fmt"
	"sync"
)

func raceCondition1Fix1() {
	ch := make(chan int)
	go func() {
		n := 0 // A local variable is only visible to one goroutine.
		n++
		ch <- n // The data leaves one goroutine...
	}()
	n := <-ch // ...and arrives safely in another.
	n++
	fmt.Println(n) // Output: 2
}

func raceCondition2Fix1() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) { // Use a local variable.
			fmt.Print(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println()
}

func raceCondition2Fix2() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		n := i // Create a unique variable for each closure.
		go func() {
			fmt.Print(n)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}

func raceCondition3Fix() {
	cntr := 0
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mux.Lock()
			cntr++ // Lock so only one goroutine at a time can access cntr.
			mux.Unlock()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			mux.Lock()
			cntr-- // Lock so only one goroutine at a time can access cntr.
			mux.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(cntr) // without WaitGroup this also fall under race condition.
}
