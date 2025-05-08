package main
import "fmt"

// Slices are dynamic 
func main() {

	//Declaring a slice --> unitialized slice is nil
	var nums[]int

	fmt.Println(nums)

	fmt.Println(nums == nil)

	fmt.Println(len(nums))

	// without nil slice
	var nums1 = make([]int, 2)
	fmt.Println(nums1)

	//capacity is the number of elements the slice can hold
	//length is the number of elements in the slice
	fmt.Println("Capacity",cap(nums1)) // 2
	fmt.Println("Length", len(nums1)) // 2

	var nums2 = make([]int, 2, 5)
	fmt.Println("Capacity",cap(nums2)) // 5
	fmt.Println("Length", len(nums2)) // 2
	fmt.Println(nums2)

	//Adding elements to the slice
	nums2 = append(nums2, 1, 2, 3, 4, 5)
	fmt.Println(nums2)
	fmt.Println("---------Capacity",cap(nums2)) // 10
	fmt.Println("Length", len(nums2)) // 7

	// nums2 have initial values of 0 at the first two indexes since we have intilized length to 2
	// So generally we keep lenth 0
	nums3 := make([]int, 0, 5)
	fmt.Println(nums3)
	fmt.Println("Capacity",cap(nums3)) // 5
	fmt.Println("Length", len(nums3)) // 0

	// another way to declare a slice
	nums4 := []int{}
	//nums := []int{1,2,3,4,5}
	fmt.Println(nums4)

	//Copy a slice
	nums5 := []int{1, 2, 3, 4, 5} 
	nums6 := make([]int, len(nums5)) 
	copy(nums6, nums5) 
	fmt.Println(nums6)

	//Slice operator
	fmt.Println(nums5[1:3]) // [2 3]
}