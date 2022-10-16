package model

type Shapes struct {
	Shapes map[string]Shape `json:"shapes"`
}

func NewShapes() *Shapes {
	return &Shapes{
		Shapes: make(map[string]Shape, 0),
	}
}
