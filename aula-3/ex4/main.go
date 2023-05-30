package main

import "fmt"

type Product struct {
	amount int
	price  float64
	id     string
}

type User struct {
	Name     string
	LastName string
	Email    string
	Products []Product
}

func (self *User) SetName(name string, lastname string) {
	self.Name = name
	self.LastName = lastname
}

func CreateUser(name string, lastname string, email string) User {
	return User{
		"Victor",
		"Vernalha",
		"a@gmail.com",
		make([]Product, 0),
	}
}

func (self *User) AddProduct(p Product) {
	self.Products = append(self.Products, p)
}

func (self *User) ClearProducts() {
	self.Products = make([]Product, 0)
}

func main() {
	user := CreateUser("Victor", "Vernalha", "a@gmail.com")
	user.AddProduct(Product{id: "MacBook", price: 20000, amount: 2})
	fmt.Println(user)
	user.ClearProducts()
	fmt.Println(user)
}
