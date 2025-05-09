package main
import "fmt"

// closures are functions that can capture variables from their surrounding context

func counter() func() int{
	count := 0
	return func() int{
		count++
		return count
	}
}

func main(){
	increament := counter()
	fmt.Println(increament()) 
	fmt.Println(increament())
}