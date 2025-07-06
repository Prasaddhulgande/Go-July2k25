package main

import (
	"fmt"
)

// When use * --> If you want to update any value in struct then need to use pointer
// If you are get any from struct then no need to use point(*) in methods

type Order struct {
	id     string
	amt    float32
	status string
	// createAt time.Time
	CM
}

type CM struct {
	name  string
	Phone string
}

// Define methodes in go
func (o *Order) changeStatus(status string) {

	o.status = status
}

func (o *Order) getAmt() float32 {
	return o.amt
}

func (o *Order) SetOrder() *Order {
	NewOrder1 := Order{
		id:     "111",
		amt:    555.00,
		status: "Ordered",
		CM: CM{
			name:  "Prasad",
			Phone: "123321",
		},
	}

	// NewCM := CM{
	// 	name:  "Prasad",
	// 	Phone: "123321",
	// }
	// fmt.Println(NewOrder1)
	return &NewOrder1

}

// Define Constructore in golang
func newOrder(id string, amt float32, status string) *Order {
	// Some Initial setup goes here...
	MyOrder := Order{
		id:     "1",
		amt:    55.00,
		status: "Received",
	}

	return &MyOrder
}

func main() {

	MyOrder := Order{
		id:     "1",
		amt:    55.00,
		status: "Received",
		// createAt: time.Now(),
	}

	fmt.Println(MyOrder)
	MyOrder.changeStatus("Conformed")
	fmt.Println(MyOrder)

	fmt.Println(MyOrder.getAmt())

	fmt.Println(newOrder("11", 30.50, "receved"))

	// Anoother way to define struct

	lang := struct {
		name   string
		isGood bool
	}{"Golang", true}

	fmt.Println(lang)

	// Create Instance of Order struct

	var NewOrder1 *Order

	// result := NewOrder1.SetOrder()
	fmt.Println(NewOrder1.SetOrder())    // Full Order
	fmt.Println(NewOrder1.SetOrder().CM) // Only CM order details

	// result.SetOrder().amt = 113.55

	fmt.Println(NewOrder1.SetOrder().name)
	// fmt.Println(NewOrder1.SetOrder().CM)

}
