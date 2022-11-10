package model

type Alternates struct {
	AbstractPartRelationships `json:"alternates"`
}

func NewAlternates() *Alternates {
	return &Alternates{
		AbstractPartRelationships: *NewAbstractRelationships(),
	}
}

func (a *Alternates) Save(filePath string) {
	Save(a, filePath)
}
