package main

import (
	"fmt"
	"os"
	"strings"
)

type StructInfo struct {
	field string
	value string
}

type Serializable interface {
	Serialize() []StructInfo
}

type Product struct {
	amount int
	price  float64
	id     string
}

func (p Product) Serialize() []StructInfo {
	return []StructInfo{
		{"id", p.id},
		{"price", fmt.Sprintf("%.2f", p.price)},
		{"amount", fmt.Sprintf("%d", p.amount)},
	}
}

func ToCSV(data []Serializable, path string) error {
	if len(data) == 0 {
		return nil
	}

	var lines strings.Builder
	for _, field := range data[0].Serialize() {
		lines.WriteString(fmt.Sprintf("%s;", field.field))
	}
	lines.WriteRune('\n')

	for _, serializable := range data {
		for _, value := range serializable.Serialize() {
			lines.WriteString(fmt.Sprintf("%s;", value.value))
		}
		lines.WriteRune('\n')
	}
	return os.WriteFile(path, []byte(lines.String()), 0644)
}

func main() {
	products := []Serializable{
		Product{10, 25.50, "Mouse"},
		Product{5, 20000, "MacBook"},
		Product{1, 10.50, "Keyring"},
	}
	ToCSV(products, "ex1/data.csv")
}
