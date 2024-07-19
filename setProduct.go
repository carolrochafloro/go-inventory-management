package main

import (
	"fmt"
)

func setProduct() {
	fmt.Print("Enter product id: ")
	var id int
	fmt.Scanln(&id)

	fmt.Print("Enter product name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter product quantity: ")
	var qtt int
	fmt.Scanln(&qtt)

	fmt.Print("Enter product price: ")
	var price float64
	fmt.Scanln(&price)

	newItem := item {
		id: id,
		name: name,
		qtt: qtt,
		price: price,
	}

	

	// file, err := os.Open("inventory.csv")

	// if err != nil {
	// 	fmt.Println("==== Error ====", err)
	// }

	// fileWriter := csv.NewWriter(file)


	
}
