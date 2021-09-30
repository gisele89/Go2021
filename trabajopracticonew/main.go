package main

import (
	"fmt"
	"trabajopractico2021.com/model"
)


func main(){
	e := model.NewStructure("AX", 3, "ABC")

	changeValues(&e, "XTZ")
	changeType(&e, "BM")
	changeLength(&e, 5)
	changeStructure(&e, "YBR", 6, "YMH")
	fmt.Println(e)
}

func changeType(e *model.Structure, typee string){
	e.Type = typee
}

func changeLength(e *model.Structure, length int){
	e.Length = length
}

func changeValues(e *model.Structure, values string){
	e.Values = values
}

func changeStructure(e *model.Structure, typee string, length int, values string){
	e.Type = typee
	e.Length = length
	e.Values = values
}
