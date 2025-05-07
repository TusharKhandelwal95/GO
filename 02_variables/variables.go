package main

import "fmt"

func main(){

	// var a string = "Hello, World!"

	//infer
	var a = "golang"
	fmt.Println(a)

	//short declaration
	b := "GO"
	fmt.Println(b)

	//if we need to only define variable without assigning value we need to use var keyword
	var c string
	fmt.Println(c)  //default value is empty string

	//var float32 = 3.14   
	//var float64 = 3.14	

	const pi = 3.14 //constant value cannot be changed
	fmt.Println(pi)

	//constant grouping
	const (
		PORT = 8080
		NAME = "localhost"
	)

	fmt.Println(PORT, NAME)

	
}