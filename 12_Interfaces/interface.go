package main

import "fmt"

type paymenter interface {
	pay(amt float32)
	refund(amt float32, acc string)
}

type payment struct {
	// gateway stripe
	// gateway razorpay
	gateway paymenter
}

func (p payment) MakePayment(amt float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amt)

	p.gateway.pay(amt)
}

func (p payment) makeRefund(amt float32, acc string) {
	p.gateway.refund(amt, acc)
}

type razorpay struct{}

func (r razorpay) pay(amt float32) {
	// Logic to make payment
	fmt.Println("Making payment using razorpay", amt)
}

type stripe struct{}

func (s stripe) pay(amt float32) {
	fmt.Println("Making payment using stripe: ", amt)
}

type fakepayment struct{}

func (f fakepayment) pay(amt float32) {
	fmt.Println("Making payment using fake gateway for testing purpose")
}

type paypal struct{}

func (p paypal) pay(amt float32) {
	fmt.Println("Using paypal payment method", 500)
}

func (p paypal) refund(amt float32, acc string) {
	fmt.Println("Successfulyy refunded amt for acc: ", amt, acc)
}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	// fakePaymentGw := fakepayment{}
	paypalPayGw := paypal{}

	newPayment := payment{
		gateway: paypalPayGw,
	}
	newPayment.MakePayment(100)
	newPayment.makeRefund(111, "99909990")
}
