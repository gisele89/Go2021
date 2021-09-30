package main

import (
	"testing"
	"entregablego.com/model"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	s := "TX04ACBD"
	type_ := "TX"
	actual, _ := model.GetStructure(s)
	assert.Equal(t, type_, actual.Type, "Las cadenas de texto resultantes son diferentes")
}

func TestLength(t *testing.T) {
	s := "TX04ACBD"
	length := 4
	actual, _ := model.GetStructure(s)
	assert.Equal(t, length, actual.Length, "La/s cadena/s de numero/s resultantes son diferentes")
}

func TestValue(t *testing.T) {
	s := "TX04ACBD"
	value := "ACBD"
	actual, _ := model.GetStructure(s)
	assert.Equal(t, value, actual.Value, "Las cadenas de texto resultantes son diferentes")
}

