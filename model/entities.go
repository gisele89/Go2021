package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

func GetStructure(s string) (Result, error) {
	if len(s) >= 4 {
		c := s[:2]
		n := s[2:4]
		v := s[4:]
		if verifyLetter(c) && verifyNumber(n) {
			number, _ := strconv.Atoi(n)
			if verifyLength(v, number) {
				return Result{c, number, v}, nil
			}
		}
	}

	return Result{}, errors.New("Estructura inválida")
}

func verifyLetter(s string) bool {
	for _, c := range s {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}
func verifyNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func verifyLength(s string, l int) bool {
	return len(s) == l
}
