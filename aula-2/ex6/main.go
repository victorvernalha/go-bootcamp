package main

import (
	"fmt"
	"time"
)

type student struct {
	Name          string
	LastName      string
	RG            string
	AdmissionDate time.Time
}

func (s student) Detail() string {
	return fmt.Sprintf("Student %s %s (RG %s) enrolled in %s", s.Name, s.LastName, s.RG, s.AdmissionDate.Format("02/01/2006"))
}

func main() {
	admissionDate, _ := time.Parse("2006-01-02", "2018-02-01")

	student := student{
		"Victor",
		"Vernalha",
		"00000000",
		admissionDate,
	}
	fmt.Println(student.Detail())
}
