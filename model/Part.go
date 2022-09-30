package model

import "strings"

// Part represents an entry of a single part in a collection with its quantity, shape, and color.
type Part struct {
	Quantity int    `json:"quantity"`
	Shape    Shape  `json:"part"`
	Color    Color  `json:"color"`
	IsSpare  bool   `json:"is_spare"`
	SetNum   string `json:"set_num"`
}

func (p *Part) LessThan(other *Part) bool {
	resultColor := strings.Compare(p.Color.Name, other.Color.Name)
	if resultColor < 0 {
		return true
	} else if resultColor > 0 {
		return false
	} else {
		resultName := strings.Compare(p.Shape.Name, other.Shape.Name)
		if resultName < 0 {
			return true
		} else {
			return false
		}
	}
}
