package main

import "fmt"

type Product interface {
	CalculateCost() float64
}

type productSize interface {
	CalculateCost(cost float64) float64
}

type productData struct {
	name string
	cost float64
}

type product struct {
	size productSize
	data productData
}

type small struct{}

type medium struct{}

type large struct{}

type Ecommerce interface {
	Total() float64
	Add(product)
}

type ecommerce struct {
	Products []product
}

const (
	Small  = 0
	Medium = 1
	Large  = 2
)

func (lp large) CalculateCost(cost float64) float64 {
	return cost*1.06 + 2500
}

func (mp medium) CalculateCost(cost float64) float64 {
	return cost * 1.03
}

func (sp small) CalculateCost(cost float64) float64 {
	return cost
}

func (p product) CalculateCost() float64 {
	return p.size.CalculateCost(p.data.cost)
}

func ProductFactory(cost float64, name string, size int) (product, error) {
	switch size {
	case Small:
		return product{size: small{}, data: productData{name, cost}}, nil
	case Medium:
		return product{size: medium{}, data: productData{name, cost}}, nil
	case Large:
		return product{size: large{}, data: productData{name, cost}}, nil
	}
	return product{}, nil
}

func (e *ecommerce) Add(p product) {
	e.Products = append(e.Products, p)
}

func (e *ecommerce) Total() float64 {
	total := 0.0
	for _, product := range e.Products {
		total += product.CalculateCost()
	}
	return total
}

func newEcommerce() ecommerce {
	return ecommerce{}
}

func main() {
	var (
		p1, _ = ProductFactory(100.0, "Mouse", Small)
		p2, _ = ProductFactory(20000.0, "Macbook", Medium)
		p3, _ = ProductFactory(5000.0, "TV", Large)
	)
	ecomm := newEcommerce()

	ecomm.Add(p1)
	ecomm.Add(p2)
	ecomm.Add(p3)

	fmt.Println(ecomm.Total())

	ecomm2 := newEcommerce()
	fmt.Println(ecomm2.Total())
}
