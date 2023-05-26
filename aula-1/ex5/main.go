package main

import "fmt"

func main() {
	var name = "Victor"

	fmt.Println(len(name))
	for _, character := range name {
		fmt.Printf("%c ", character)
	}
	fmt.Println()
}