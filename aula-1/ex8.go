package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var target = "Benjamin"
	fmt.Println("Idade de", target, "é", employees[target])

	var over21_count = 0
	for _, age := range employees {
		if age > 21 {
			over21_count++
		}
	}

	fmt.Println("Employees over 21:", over21_count)

	employees["Frederico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}