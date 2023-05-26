package main

import (
	"fmt"
)

const (
	A = 0
	B = 1
	C = 2
)

func CalculateSalary(category int, workedHours int) (salary float64) {
	baseSalary := []int{3000, 1500, 1000}
	multipliers := []float64{1.5, 1.2, 1}

	salary = float64(baseSalary[category] * workedHours)
	if workedHours > 160 {
		salary *= multipliers[category]
	}
	return
}

func main() {
	fmt.Println(CalculateSalary(C, 162))
	fmt.Println(CalculateSalary(B, 176))
	fmt.Println(CalculateSalary(A, 172))
}
