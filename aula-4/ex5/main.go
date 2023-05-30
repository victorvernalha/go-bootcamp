package main

import (
	"fmt"
	"os"
)

type Customer struct{}

func ParseCustomers(filePath string) (customers []Customer, err error) {
	_, err = os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("error opening file %s", filePath))
	}

	// Continue processing customers...
	return
}

func main() {
	ParseCustomers("customers.txt")
}
