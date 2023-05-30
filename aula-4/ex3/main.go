package main

import (
	"fmt"
)

func main() {
	var salary = 11000
	var minimumValue = 15000

	if salary < minimumValue {
		err := fmt.Errorf("%d is below minimum value %d", salary, minimumValue)
		fmt.Print(err)
	} else {
		fmt.Println("Needs to pay taxes")
	}
}
