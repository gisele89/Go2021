package main

import (
	"fmt"

	"entregablego.com/model"
)

func main() {
	result, err := model.GetStructure("TX04ACBD")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
