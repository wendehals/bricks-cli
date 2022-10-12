package model

type Shapes struct {
	Shapes []Shape `json:"shapes"`
}

func NewShapes() *Shapes {
	return &Shapes{
		Shapes: []Shape{},
	}
}
