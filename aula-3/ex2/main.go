package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	amount int
	price  float64
	id     string
}

func ReadProductsFromCSV(path string) ([]Product, error) {
	products := []Product{}

	fileContents, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	for index, line := range strings.Split(string(fileContents), "\n") {
		if index == 0 {
			continue
		}

		split := strings.Split(line, ";")
		if len(split) < 3 {
			continue
		}

		fmt.Println(split)
		id := split[0]
		price, _ := strconv.ParseFloat(split[1], 64)
		amount, _ := strconv.Atoi(split[2])

		products = append(products, Product{amount, price, id})
	}

	return products, nil
}

func DisplayProducts(products []Product) {
	fmt.Printf("ID\tAmount\tPrice\n")
	for _, product := range products {
		fmt.Printf("%s\t%d\t%.2f\n", product.id, product.amount, product.price)
	}
}

func main() {
	products, _ := ReadProductsFromCSV("ex2/data.csv")
	DisplayProducts(products)
}
