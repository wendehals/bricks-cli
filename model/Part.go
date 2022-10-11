package model

// Part represents an entry of a single part in a collection with its quantity, shape, and color.
type Part struct {
	Quantity int    `json:"quantity"`
	Shape    Shape  `json:"part"`
	Color    Color  `json:"color"`
	IsSpare  bool   `json:"is_spare"`
	SetNum   string `json:"set_num"`
}

func (p *Part) Compare(other *Part) int {
	resultColor := p.Color.Compare(&other.Color)
	if resultColor == 0 {
		return p.Shape.Compare(&other.Shape)
	}
	return resultColor
}

func (p *Part) CompareByQuantity(other *Part) int {
	resultColor := 0
	if p.Quantity > other.Quantity {
		resultColor = 1
	} else if p.Quantity < other.Quantity {
		resultColor = -1
	}

	if resultColor == 0 {
		return p.Shape.Compare(&other.Shape)
	}
	return resultColor
}
