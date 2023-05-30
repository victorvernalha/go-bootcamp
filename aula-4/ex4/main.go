package main

import (
	"errors"
	"fmt"
)

func CalculateSalary(hoursWorked, hourlyWage int) (total float64, err error) {
	if hoursWorked < 80 {
		err = errors.New("error: employee can't have worked less than 80 hours in a week")
	}

	total = float64(hoursWorked * hourlyWage)

	const taxCutoff = 20000
	if total >= taxCutoff {
		total *= 0.9
	}

	const minValue = 15000
	if total <= minValue {
		err = fmt.Errorf("%w\nerror: salary %.2f is below minimum value %d", err, total, minValue)
	}

	return
}

func main() {
	var hourlyWage = 1
	var hoursWorked = 90

	salary, err := CalculateSalary(hoursWorked, hourlyWage)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}
}
