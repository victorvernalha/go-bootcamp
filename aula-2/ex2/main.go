package main

import (
	"fmt"
)

func Mean(first_grade float64, grades ...float64) float64 {
	total := first_grade
	for _, grade := range grades {
		total += grade
	}

	return total / float64(len(grades)+1)
}

func main() {
	fmt.Println(Mean(1, 2, 3, 4, 5))
	fmt.Println(Mean(10))
}
