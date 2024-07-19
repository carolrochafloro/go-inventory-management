package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

	var newItem = []string {
		strconv.Itoa(id), 
		name,
		strconv.Itoa(qtt),
		strconv.FormatFloat(price, 'f', 2, 64),
	}

	writeProduct(newItem)
}

func writeProduct(s []string) {

	file, err := os.OpenFile("inventory.csv", os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("==== Error ====", err)
	}

	defer file.Close()

	fileWriter := csv.NewWriter(file)
	defer fileWriter.Flush()
	fileWriter.Write(s)

	if err := fileWriter.Error(); err != nil {
        fmt.Println("==== Error ====", err)
    }

}