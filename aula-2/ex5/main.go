package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogcalculation(amount int) int {
	return amount * 10000
}

func catcalculation(amount int) int {
	return amount * 5000
}

func hamstercalculation(amount int) int {
	return amount * 250
}

func tarantulacalculation(amount int) int {
	return amount * 150
}

func Animal(animal string) (func(int) int, string) {
	switch animal {
	case dog:
		return dogcalculation, ""
	case cat:
		return catcalculation, ""
	case hamster:
		return hamstercalculation, ""
	case tarantula:
		return tarantulacalculation, ""
	}
	return nil, "Animal not accounted for"
}

func main() {
	animalDog, _ := Animal(dog)
	animalCat, _ := Animal(cat)

	var amount int
	amount += animalDog(5)
	amount += animalCat(8)
	fmt.Println(amount)
}
