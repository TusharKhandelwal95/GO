package main
import "fmt"

//We use slices more than arrays in go, slices are similar to vectors

func main(){
	// Array intialization
	var arr[5]int

	// Finding length of array
	fmt.Println("Lenght of array is ", len(arr))

	// Assigning value to first element of array
	arr[0] = 1
	fmt.Println("First element of array is ", arr[0])

	// Printing whole array
	fmt.Println("Array is ", arr)

	// Array with initial values
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array is ", nums)

	//2-d array
	twoD := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println("2-D array is ", twoD)


}