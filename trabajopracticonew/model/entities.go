package model

type Structure struct{
	Type string
	Length int
	Values string
}

func NewStructure(typee string, length int, values string) Structure{
	return Structure{typee, length, values}
}