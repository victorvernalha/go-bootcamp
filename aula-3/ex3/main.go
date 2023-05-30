package main

import "fmt"

type User struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func (self *User) SetName(name string, lastname string) {
	self.Name = name
	self.LastName = lastname
}

func (self *User) SetAge(age int) {
	self.Age = age
}

func (self *User) SetEmail(email string) {
	self.Email = email
}

func (self *User) SetPassword(password string) {
	self.Password = password
}

func main() {
	user := User{
		"Victor",
		"Vernalha",
		22,
		"aaaa@email.com",
		"1234",
	}
	user.SetAge(23)
	fmt.Println(user)
}
