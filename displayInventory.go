package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func displayInventory(s string) {

	// read file
	file, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("==== Error ====", err)
		os.Exit(1)
	}

	defer file.Close()

	if err != nil {
		fmt.Println("==== Error ====", err)
		os.Exit(1)
	}

	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Println("==== Error ====", err)
		os.Exit(1)
	}

	if fileInfo.Size() == 0 {
		setNewFile(file)
	}

	fileReader := csv.NewReader(file)

	content, err2 := fileReader.ReadAll()

	if err2 != nil {
		fmt.Println("==== Error ====", err2)
	}

	printContent := convertSliceToStruct(content)

	for _, item := range printContent {
		fmt.Printf("Id: %d Name: %s Qtt: %d Price: %.2f\n", item.id, item.name, item.qtt, item.price)
	}

}

func setNewFile(f *os.File) {

	header := []string {
		"ID",
		"Name",
		"Qtt",
		"Price",
	}

	fileWriter := csv.NewWriter(f)

	defer fileWriter.Flush()
	fileWriter.Write(header)
}


func convertSliceToStruct(content [][]string) []item {
	var items []item 

	for i, line := range content {

		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(line[0])

		if err != nil {
			fmt.Println("Error converting ID:", err)
			continue
		}

		qtt, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Println("Error converting Quantity:", err)
			continue
		}

		price, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			fmt.Println("Error converting Price:", err)
			continue
		}

		item := item{
			id:    id,
			name:  line[1],
			qtt:   qtt,
			price: price,
		}

		items = append(items, item)

	}

	return items
}