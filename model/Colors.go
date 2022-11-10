package model

type Colors struct {
	Colors []Color `json:"colors"`
}

func NewColors() *Colors {
	return &Colors{
		Colors: []Color{},
	}
}

func (c *Colors) Save(filePath string) {
	Save(c, filePath)
}
