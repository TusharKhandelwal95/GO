package main
import "fmt"

//Variadic functions are functions that can take a variable number of arguments
// example: Println(a, b, c)

// ...int means that the function can take a variable number of int arguments
func sum(nums ...int)int{
	total := 0
	for _, a := range nums{
		total += a
	}
	return total
}

// we can use interface{} to accept any type of argument
func printAll(args ...interface{}){
	for _,arg := range args{
		fmt.Println(arg)
	}
}

func main(){
	fmt.Println("Hello", "World", 1, 2, 3)

	fmt.Println(sum(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(nums...)) // use ... to unpack the slice
	printAll("Hello", 1, 2.5, true, []int{1, 2, 3})
}