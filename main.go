package main

import (
	"fmt"
	"os"
)

func main() {
	displayInventory("inventory.csv")

	c := menu()

	if c == "y" {
		setProduct()
		displayInventory("inventory.csv")
		menu()
	}

	os.Exit(0)

}
func menu() string {
	fmt.Println("Do you want to create a new product? y/n ")
	var c string
	fmt.Scanln(&c)

	if c != "y" && c != "n" {
		fmt.Println("You must type y for yes or n for no")
		menu()
	}
	
	return c
}