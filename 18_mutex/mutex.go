/*
Mutex - It means mutually exclusion used to prevent race condition
      - It locks the resource when one process is modifying it
	  - One drawback is the lock becomes the bottlenecks like are mutliple go routines have to wait for the resource now
*/

package main

import (
	"fmt"
	"sync"
)

// type post struct{
// 	views int
// }

// func (p*post) inc(wg* sync.WaitGroup){
// 	defer wg.Done()
// 	p.views += 1
// }

// func main(){
// 	var wg sync.WaitGroup
// 	myPost := post{views:0}

// 	for i:=0; i<100; i++{
// 		wg.Add(1)
// 		go myPost.inc(&wg) //running concurrently
// 	} 
// 	wg.Wait()
// 	fmt.Println(myPost.views) // we are getting random value because all processes are modify together that occur race condition

// }

//---------------------------------------------------

type post struct{
	views int
	mu sync.Mutex // good practice to declare in the struct where it is used
}

func (p*post) inc(wg* sync.WaitGroup){
	defer func(){
		p.mu.Unlock()  // good practice to keep unlock in defer
		wg.Done()
	}()
	p.mu.Lock()  // lock at the place where modification is done not on entire function in start
	p.views += 1
	
}

func main(){
	var wg sync.WaitGroup
	myPost := post{views:0}

	for i:=0; i<100; i++{
		wg.Add(1)
		go myPost.inc(&wg) //running concurrently
	} 
	wg.Wait()
	fmt.Println(myPost.views) // 100 - now no error since mutex is running
}