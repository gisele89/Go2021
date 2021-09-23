package main

import (
	"fmt"

	"entregablego.com/model"
)

func main() {
	result, err := model.GetStructure("TX04ABCD")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
