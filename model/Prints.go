package model

type Prints struct {
	AbstractPartRelationships `json:"prints"`
}

func NewPrints() *Prints {
	return &Prints{
		AbstractPartRelationships: *NewAbstractRelationships(),
	}
}
