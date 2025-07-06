package main

import "fmt"

// Define enums using type & const key
type OrderStatus int

const (
	Received OrderStatus = iota  // iota is a special type of integer
	Conformed
	Prepared
	Delivered
)

func chnageOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	chnageOrderStatus(Prepared)
}
