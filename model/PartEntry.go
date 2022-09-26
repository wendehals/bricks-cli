package model

import "strings"

// PartEntry represents an entry of a single part in a collection with its quantity and color.
type PartEntry struct {
	Quantity int    `json:"quantity"`
	Part     Part   `json:"part"`
	Color    Color  `json:"color"`
	IsSpare  bool   `json:"is_spare"`
	SetNum   string `json:"set_num"`
}

func (p *PartEntry) LessThan(other *PartEntry) bool {
	resultColor := strings.Compare(p.Color.Name, other.Color.Name)
	if resultColor < 0 {
		return true
	} else if resultColor > 0 {
		return false
	} else {
		resultName := strings.Compare(p.Part.Name, other.Part.Name)
		if resultName < 0 {
			return true
		} else {
			return false
		}
	}
}
