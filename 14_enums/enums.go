package main
import "fmt"

//enumerated types

// type orderStatus int
type orderStatus string

const (
	// each constant is assigned a value starting from 0
	// Received orderStatus = iota
	// Confirmed
	// Prepared
	// Shipped
	// Delivered  

	Received orderStatus = "Received"
	Confirmed = "Confirmed"
	Prepared = "Prepared"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("Order status changed to", status)
}

func main() {
	changeOrderStatus(Confirmed) //Order status changed to 1
}	