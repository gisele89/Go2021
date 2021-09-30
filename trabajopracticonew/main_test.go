package main

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"trabajopractico2021.com/model"
  )

func TestValues(t *testing.T){
	e := model.NewStructure("AX", 3, "ABC")

	changeStructure(&e, "XY", 5, "WXY")
	changeLength(&e, 4)
	changeType(&e, "AR")
	changeValues(&e, "FGH")

	// FALTA AGREGAR CONDICIOENS PARA VALIDAR LOS ASSERTS


	
}
