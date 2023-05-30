package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary = 11000
	var minimumValue = 15000

	if salary < minimumValue {
		err := errors.New(fmt.Sprintf("%d is below minimum value %d", salary, minimumValue))
		fmt.Print(err)
	} else {
		fmt.Println("Needs to pay taxes")
	}
}
