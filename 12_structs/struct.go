package main
import ("fmt"; "time")

//like classes in C++

type Customer struct{
	name string 
	age int
}

type Order struct{
	id int
	amt float32
	status string
	createdAt time.Time // nanosecond precision

	Customer // *****embedding struct****
}

// there is no constructor in go
func newOrder(id int, amt float32, status string) *Order{
	myOrder := Order{  // pointer is returned
		id: id,
		amt: amt,
		status: status,
	}
	return &myOrder  
}

// attaching methods to struct

// func (o Order) changeStatus(newStatus string){  // by convention, the first letter of struct field as o in (o Order)
// 	o.status = newStatus // here a copy of struct is passed 
// }

func (o *Order) changeStatus(newStatus string){  //passing by refrence
	o.status = newStatus  // struct derefrence 
}

func main(){

	//Creating instance of struct
	// if you don't set any value, it will set to zero value
	myOrder := Order{
		id: 1,
		amt: 50.20,
		status: "received",
	}

	// fmt.Println("Order struct:", myOrder) //Order struct: {1 50.2 received {0 0 <nil>}}

	// //Adding field after creating instance
	// myOrder.createdAt = time.Now()

	// fmt.Println("Order struct:", myOrder) //Order struct: {1 50.2 received {13979069316808478424 1 0x8239e0}}

	//Accessing fields
	fmt.Println("Status:" ,myOrder.status) 

	myOrder.changeStatus("confirmed") 

	fmt.Println("Status:" ,myOrder.status) 


	myOrder2 := newOrder(2, 100.50, "pending") 
	fmt.Println("Order 2", myOrder2) 


	// if we want to use struct only once, we can use anonymous struct
	myLang := struct{
		lang string
		isGood bool
	}{"Go", true}

	fmt.Println("Anonymous struct:", myLang) 
}