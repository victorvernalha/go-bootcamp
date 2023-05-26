package main

import "fmt"

func CalculateTax(salary float64) int {
	if salary < 150000 {
		return 17
	}
	return 27
}

func main() {
	fmt.Println(CalculateTax(150000))
	fmt.Println(CalculateTax(50000))
}
