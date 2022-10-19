package model

import "github.com/wendehals/bricks/utils"

type Shapes struct {
	Shapes map[string]Shape `json:"shapes"`
}

func NewShapes() *Shapes {
	return &Shapes{
		Shapes: make(map[string]Shape, 0),
	}
}

func (s *Shapes) Save() {
	Save(s, utils.ShapesPath())
}
