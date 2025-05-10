package main
import "fmt"

// interface is a contract that defines a set of methods that a type must implement
// In go we dont have to explicitly declare that a type implements an interface

type paymenter interface {
	pay(amount float32)
}

// type payment struct{}
type payment struct{
	// gateway stripe
	gateway paymenter
}


// Here we are voilating the OPEN CLOSE principle
// which states that the software entities classes, modules, functions should be open for extension but closed for modification.
func (p payment) makePayment(amount float32){
	// makePaymentGw := razorpay{}
	// makePaymentGw.pay(amount)

	// makePaymentGw := stripe{}
	// makePaymentGw.pay(amount)  

	//OR

	p.gateway.pay(amount)


}

type razorpay struct {}
func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Payment made using Razorpay of", amount)
}

// suppose now we want to change payment to use stripe instead of razorpay
type stripe struct {}
func (s stripe) pay(amount float32){
	// logic to make payment
	fmt.Println("Payment made using Stripe of", amount)
}

func main(){
	// p := payment{}

	// stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}
	p := payment{
		// gateway: stripePaymentGw,
		gateway: razorpayPaymentGw,
	}
	p.makePayment(1000.0)
}