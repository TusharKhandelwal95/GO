package main
import "fmt"
// generics were introduced in go 1.18
// generics allow us to write functions and data structures that can work with any type

func printSlice(items []int){
	for _,item := range items{
		fmt.Println(item)
	}
}

// func printSliceGeneric[T any](items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// if we want to limit the type
// func printSliceGeneric[T int | string](items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// another way to write the above function using interface{} as a type parameter
func printSliceGeneric[T interface{}](items []T){
	for _,item := range items{
		fmt.Println(item)
	}
}

// we can also use generics with structs
type genericStruct[T any, V string | bool] struct {
	name T
	age V
}

func main() {
	nums := []int{1,2,3}
	names := []string{"Go", "C++", "Python"}

	printSlice(nums)
	// printSlice(names) // this will not work because the function is not generic

	printSliceGeneric(nums)
	printSliceGeneric(names)

	// generic struct
	g := genericStruct[string, bool]{ "Go", true}
	
	fmt.Println(g.name)
	fmt.Println(g.age)

}