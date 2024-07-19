package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

	fmt.Println(content)

}

func setNewFile(f *os.File) {

	header := []string {
		"Id",
		"Name",
		"Qtt",
		"Price",
	}

	fileWriter := csv.NewWriter(f)

	defer fileWriter.Flush()
	fileWriter.Write(header)
}
