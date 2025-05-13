// Communication between goroutines using channels

package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "sync"
)

//INTRODUCTION
// func main(){
// 	num := make(chan int)
// 	num<-1
// 	fmt.Println(num) //fatal error: all goroutines are asleep - deadlock!

// }


//REMOVE ERROR
// func processNum(num chan int){
// 	fmt.Println("received: ",<-num)
// }

// func main(){
// 	num := make(chan int)

// 	go processNum(num)
// 	num<-5

// 	time.Sleep(time.Second*2)

// }
//---------------------------------------------

//SENDING VALUE INFINITELY
// func processNum(numChan chan int){
// 	for num := range numChan {
// 		fmt.Println("received: ",num)
// 		time.Sleep(time.Second*2)
// 	}
// }

// func main(){
// 	numChan := make(chan int)

// 	go processNum(numChan)
	
// 	for{
// 		numChan <- rand.Intn(100)
// 	}

// }
//---------------------------------------------------

//RECEIVING
// func sum(n1 int, n2 int, numChan chan int){
// 	result := n1 + n2
// 	numChan <- result
// }

// func main(){
// 	result := make(chan int)

// 	go sum(5,6,result)
	
// 	fmt.Println(<-result) // 11

// }
//-------------------------------------------------------



//USING DEFER CHANNEL AS WAITGROUP -- GOROUTINE SYNCHRONIZER
// func task(done chan bool){
// 	defer func(){done <- true}() //using defer make it run after the complete outer function is done
// 	fmt.Println("Processing....")
// }

// func main(){
// 	done := make(chan bool)
// 	go task(done)
// 	<-done
// }
//-----------------------------------------------------------


//BUFFERED CHANNEL - WE CAN SEND LIMITED DATA WITHOUT BLOCKING
// func main(){
// 	emailChan := make(chan string, 100) //buffered channel

// 	emailChan <- "1@example.com"
// 	emailChan <- "2@example.com"

// 	//Now we can reach here since we can send 100 therefore it will be non blocking
// 	fmt.Println(<-emailChan)
// 	fmt.Println(<-emailChan)

// }
//------------------------------------------------------


//CLOSING CHANNEL
// func emailSender(emailChan chan string, done chan bool){

// 	defer func(){done<-true}()

// 	//this for loop in never ending and thus defer above will never work
// 	for email := range emailChan{
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}

// }

// func main(){
// 	emailChan := make(chan string, 100) //buffered channel
// 	done := make(chan bool)

// 	go emailSender(emailChan, done)
	
// 	for i:=0; i<10; i++{
// 		emailChan <- fmt.Sprintf("%d@gmail.com",i)
// 	}
	
// 	//we close the channel after work done to prevent Deadlock
// 	close(emailChan) //IMP TO CLOSE

// 	<-done
// 	// Deadlock will occur after printing 9@gmail.com 

// }
//-------------------------------------------------------------------

//MULTIPLE CHANNEL

// func main(){
// 	chan1 := make(chan int)
// 	chan2 := make(chan string)

// 	go func(){
// 		chan1 <- 10
// 	}()

// 	go func(){
// 		chan2 <- "pong"
// 	}()

// 	for i:=0; i<2; i++{
// 		select{
// 		case chan1Val := <-chan1:
// 			fmt.Println("Accesing value from chan1 ",chan1Val)
// 		case chan2Val := <-chan2:
// 			fmt.Println("Accesing value from chan2 ",chan2Val)
// 		}
// 	}
// }
//-----------------------------------------------------------------


//WE CAN MAKE A CHANNEL TO ONLY RECEIVE/SEND ONLY CHANNEL
func emailSender(emailChan <-chan string, done chan<- bool){  

	//email chan -- receive only channel
	//done chan -- send only channel

	defer func(){done<-true}()

	//this for loop in never ending and thus defer above will never work
	for email := range emailChan{
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}

}


