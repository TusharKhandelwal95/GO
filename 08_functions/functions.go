package main
import "fmt"

// Functions are first-class citizens in Go
// Functions can be assigned to variables, passed as arguments, and returned from other functions

// func add(a int, b int)int{
// 	return a + b
// }

func add(a, b int)int{
	return a + b
}

func getLanguage()(string, string, string){
	return "Go", "C++", "Python"
}

func processIt(fn func(a int)int){
	fn(10)
}

func main(){
	fmt.Println(add(1,2))

	lang1, lang2, _ := getLanguage()
	fmt.Println(lang1, lang2)

	fn := func(a int)int{
		fmt.Println("Anonymous function")
		return a * 2
	}

	processIt(fn)
	
}