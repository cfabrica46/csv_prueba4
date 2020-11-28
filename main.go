package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type animal struct {
	name  string
	color string
	bioma string
}

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

	animals := []animal{}

	for i1 := range data {

		animal := animal{name: data[i1][0], color: data[i1][1], bioma: data[i1][2]}

		animals = append(animals, animal)
	}

	fmt.Println(animals)
}
