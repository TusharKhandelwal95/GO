package main

import (
	"fmt"
	"sync"
)

// Routines are lightweight threads managed by the Go runtime.
// They are cheaper than OS threads and can be created in large numbers.

// func task(id int){
// 	fmt.Println("Doing task", id)
// }

func task(id int, w* sync.WaitGroup){
	defer w.Done() // Decrement the wait group counter when the goroutine is done
	fmt.Println("Doing task", id)
}

func main(){
	
	// Goroutines will not run in order
	// The Go runtime will schedule them as it sees fit

	// we use wait groups to wait for all goroutines to finish
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		wg.Add(1) // Add a task to the wait group
		go task(i, &wg) // Start a new goroutine
		// Nothing will be printed here because the main function will exit before the goroutines finish

		// go func(i int){
		// 	fmt.Println(i)
		// }(i)
	}

	// time.Sleep(1 * time.Second) // Wait for goroutines to finish -- generally we dont do like this because we dont know how long they will take

	wg.Wait() // Wait for all goroutines to finish


	fmt.Println("All tasks done")

}