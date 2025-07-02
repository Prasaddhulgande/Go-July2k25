package main

import (
	"fmt"
	"reflect"
)

func main() {

	M := make(map[string]string)

	M["Name"] = "Prasad"
	M["pin"] = "431720"
	M["state"] = "MH"

	fmt.Println(M)
	fmt.Println(M["pin"])
	fmt.Println(len(M))

	delete(M, "Name")
	fmt.Println(M)

	// 2nd type to define map

	m := map[string]int{"ID": 111, "Phone": 911210, "pin": 2143062}

	fmt.Println(m)

	val, boolean := m["ID"]

	fmt.Println(boolean)
	fmt.Println(val)

	m1 := map[string]int{"ID": 111, "Phone": 911210, "pin": 2143062}

	fmt.Println(reflect.DeepEqual(m, m1))
}
