package main

import "fmt"

type Product struct {
	id     string
	amount int
	price  float64
}

type Service struct {
	id            string
	price         float64
	minutesWorked int
}

type Maintenance struct {
	id    string
	price float64
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func CalculateProductsCost(products []Product, c chan float64) {
	total := 0.0
	for _, product := range products {
		total += product.price * float64(product.amount)
	}
	c <- total
}

func CalculateServicesCost(services []Service, c chan float64) {
	total := 0.0
	for _, service := range services {
		halfHoursWorked := Max(1, service.minutesWorked/30)
		total += service.price * float64(halfHoursWorked)
	}
	c <- total
}

func CalculateMaintenancesCost(maintenances []Maintenance, c chan float64) {
	total := 0.0
	for _, maintenance := range maintenances {
		total += maintenance.price
	}
	c <- total
}

func main() {
	products := []Product{
		{"A", 3, 10},
		{"B", 1, 50},
	}
	services := []Service{
		{"A", 10, 20},
		{"B", 20, 70},
		{"C", 100, 60},
	}
	maintenances := []Maintenance{}

	c := make(chan float64)
	total := 0.0

	go CalculateMaintenancesCost(maintenances, c)
	go CalculateProductsCost(products, c)
	go CalculateServicesCost(services, c)

	for i := 0; i < 3; i++ {
		total += <-c
	}

	fmt.Println(total)
}
