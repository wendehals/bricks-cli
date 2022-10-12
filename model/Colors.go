package model

type Colors struct {
	Colors []Color `json:"colors"`
}

func NewColors() *Colors {
	return &Colors{
		Colors: []Color{},
	}
}
