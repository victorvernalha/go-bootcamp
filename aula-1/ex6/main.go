package main

import "fmt"

func main() {
	const (
		age                 = 22
		employed            = true
		years_of_experience = 3
		salary              = 8000
	)

	if age > 22 && employed && years_of_experience > 1 {
		if salary > 100 {
			fmt.Println("No interest")
		} else {
			fmt.Println("Interest")
		}
	} else {
		fmt.Println("No line of credit")
	}
}
