package main
import "fmt"

func main() {

	for i := 0; i<5; i++{
		fmt.Println(i)
	}

	//while loop
	j := 0
	for j<5{
		fmt.Println(j)
		j++
	}

	//infinite loop
	k:=0
	for{
		if k>5{
			fmt.Println("k grater than 5")
			break
		}
		fmt.Println(k)
	}
}