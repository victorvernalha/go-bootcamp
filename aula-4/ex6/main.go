package main

import (
	"errors"
	"fmt"
	"os"
)

type Customer struct {
	Id      int
	Name    string
	RG      string
	Phone   string
	Address string
}

type ErrCustomerAlreadyExists struct {
	CustomerRG string
}

func (err ErrCustomerAlreadyExists) Error() string {
	return fmt.Sprintf("customer with RG %s already exists", err.CustomerRG)
}

func ParseCustomers(filePath string) (customers []Customer, err error) {
	defer func() {
		recovery := recover()
		if err := recovery.(error); errors.Is(err, os.ErrNotExist) {
			fmt.Println("error: file not found. Continuing anyway...")
		} else if err != nil {
			panic(err)
		}
	}()

	_, err = os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("error opening file %s: %w", filePath, err))
	}

	// Continue processing customers...
	return
}

func GenerateUniqueIDFor(RG string, customers []Customer) (int, error) {
	for _, customer := range customers {
		if customer.RG == RG {
			err := ErrCustomerAlreadyExists{RG}
			return 0, err
		}
	}
	return len(customers), nil
}

func AddCustomer(customers []Customer, name, RG, phone, address string) []Customer {
	id, err := GenerateUniqueIDFor(RG, customers)
	if err != nil {
		panic(err)
	}

	customer := Customer{
		id, name, RG, phone, address,
	}
	return append(customers, customer)
}

func main() {
	customers, _ := ParseCustomers("customers.txt")
	customers = AddCustomer(customers, "Victor", "1234", "55...", "Py Street")
	customers = AddCustomer(customers, "João", "4321", "55...", "TS Street")
	customers = AddCustomer(customers, "Tiago", "9999", "55...", "C# Street")
	customers = AddCustomer(customers, "Victor", "1234", "55...", "Py Street")

	fmt.Println(customers)
}
