package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func displayInventory(s string) {

	// read file
	file, err := os.Open(s)

	if err != nil {
		fmt.Println("==== Error ====", err)
		os.Exit(1)
	}

	fileReader := csv.NewReader(file)

	content, err2 := fileReader.ReadAll()

	if err2 != nil {
		fmt.Println("==== Error ====", err2)
	}

	fmt.Println(content)

}