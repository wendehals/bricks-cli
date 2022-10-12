package model

type Molds struct {
	AbstractPartRelationships `json:"molds"`
}

func NewMolds() *Molds {
	return &Molds{
		AbstractPartRelationships: *NewAbstractRelationships(),
	}
}
