package main

import "fmt"

type SalaryError struct {
	Salary   int
	MinValue int
}

func (self SalaryError) Error() string {
	return fmt.Sprintf("%d is below minimum value %d", self.Salary, self.MinValue)
}

func main() {
	var salary = 11000
	var minimumValue = 15000

	if salary < minimumValue {
		err := SalaryError{salary, minimumValue}
		fmt.Print(err)
	} else {
		fmt.Println("Needs to pay taxes")
	}
}
