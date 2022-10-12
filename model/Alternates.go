package model

type Alternates struct {
	AbstractPartRelationships `json:"alternates"`
}

func NewAlternates() *Alternates {
	return &Alternates{
		AbstractPartRelationships: *NewAbstractRelationships(),
	}
}
