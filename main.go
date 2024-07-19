package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	// read file
	f, err := os.Open("inventory.csv")

	if err != nil {
		fmt.Println("==== Error ====", err)
		os.Exit(1)
	}

	fileReader := csv.NewReader(f)

	content, err2 :=fileReader.Read()

	if err2 != nil {
		fmt.Println("==== Error ====", err2)
	}

	fmt.Println(content)

	// file implements Reader interface
	// io.Copy(os.Stdout, f)

}