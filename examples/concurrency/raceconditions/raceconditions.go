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

func raceCondition1() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n++ // read, increment, write
		close(wait)
	}()
	n++ // conflicting access
	<-wait
	fmt.Println(n) // Output: <unspecified>
}

func raceCondition2() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i) // Output: <unspecified> but mostly 5555
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()
}
