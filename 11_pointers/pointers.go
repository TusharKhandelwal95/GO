package main
import "fmt"


func changeNum(num *int){
	*num = 10 // derefrencing
	fmt.Println("num inside changeNum:", *num)
}

func main(){

	num := 1
	fmt.Println("num before chnageNum:", num)
	fmt.Println("address of num:", &num) 
	changeNum(&num) 
	fmt.Println("num after chnageNum:", num)
}