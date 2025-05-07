package main
import "fmt"

func main() {
	// if else statement
	age := 18
	if age < 18 {
		fmt.Println("minor")
	} else if age == 18 {
		fmt.Println("exactly 18")
	} else {
		fmt.Println("adult")
	}

	// varaible declaration in if statement
	// if age := 20; age < 18 {
	// 	fmt.Println("minor")
	// } else{
	// 	fmt.Println("adult")
	// }

	// switch statement
	x := 2
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	default:
		fmt.Println("x is greater than 3")
	}

	// multiple condition switch
	switch x {
	case 1, 2:
		fmt.Println("x is 1 or 2")
	case 3, 4:
		fmt.Println("x is 3 or 4")
	default:
		fmt.Println("x is greater than 4")
	}

	// type switch

	whoAmI := func(i interface{}){
		switch i.(type) {
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I am something else")
		}
	}

	whoAmI("golang")

}

/* 
exactly 18
x is 2
x is 1 or 2
I am a string
*/