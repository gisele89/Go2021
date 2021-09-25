package main

import (
	"testing"
	"entregablego.com/model"
	"github.com/stretchr/testify/assert"
)

func TestGetStructure(t *testing.T) {
	s := "TX04ACBD"
	expected := "TX 4 ACBD"
	actual, _ := model.GetStructure(s)
	assert.Equal(t, expected, actual, "Las cadenas de texto resultantes son diferentes")
}

/*func TestverifyLetter(t *testing.T) {
	s := "TX04ACBD"
	expected := true
	actual, _ := model.verifyLetter(s)
	assert.Equal(t, actual, expected, "Las cadenas de texto resultantes son diferentes")
}

func verifyNumber(t *testing.T) {

}

func verifyLength(t *testing.T) {

}
*/