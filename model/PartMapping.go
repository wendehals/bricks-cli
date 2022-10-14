package model

type PartMapping struct {
	Quantity    int    `json:"quantity"`
	Original    Part   `json:"original"`
	Substitutes []Part `json:"substitutes"`
}

func NewPartMapping() *PartMapping {
	return &PartMapping{
		Original:    Part{},
		Substitutes: []Part{},
	}
}

func (p *PartMapping) Compare(other *PartMapping) int {
	return p.Original.CompareByColorAndName(&other.Original)
}
