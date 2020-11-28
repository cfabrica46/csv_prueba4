package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	archivo, err := os.OpenFile("databases.csv", os.O_RDWR, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	lector := csv.NewReader(archivo)

	data, err := lector.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
