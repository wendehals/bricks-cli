package model

// Part represents an entry of a single part in a collection with its quantity, shape, and color.
type Part struct {
	Quantity int    `json:"quantity"`
	Shape    Shape  `json:"part"`
	Color    Color  `json:"color"`
	IsSpare  bool   `json:"is_spare"`
	SetNum   string `json:"set_num"`
}

func NewPart() *Part {
	return &Part{
		Shape: Shape{},
		Color: Color{},
	}
}

func (p *Part) CompareByColorAndName(other *Part) int {
	result := p.Color.Compare(&other.Color)
	if result == 0 {
		return p.Shape.Compare(&other.Shape)
	}
	return result
}

func (p *Part) CompareByQuantityAndName(other *Part) int {
	result := 0
	if p.Quantity > other.Quantity {
		result = 1
	} else if p.Quantity < other.Quantity {
		result = -1
	}

	if result == 0 {
		return p.Shape.Compare(&other.Shape)
	}
	return result
}
