package main

import (
	"fmt"

	"github.com/ArtemAleynikov/models"
)

type data []*models.Person

func main() {
	persons := data{
		{Age: 12, Name: "Artem"},
	}

	fmt.Println(*persons[0])
}
