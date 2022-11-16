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

func (s *Shapes) GetShape(number string) (Shape, bool) {
	shape, found := s.Shapes[number]
	return shape, found
}

func (s *Shapes) GetImageURL(number string) string {
	if shape, found := s.Shapes[number]; found {
		if shape.ImageURL != "" {
			return shape.ImageURL
		}
	}

	return ""
}

func (s *Shapes) ComplementShapesData(parts []Part) {
	for _, part := range parts {
		if shape, found := s.Shapes[part.Shape.Number]; found {
			if shape.ImageURL == "" {
				shape.ImageURL = part.Shape.ImageURL
				s.Shapes[part.Shape.Number] = shape
			}
		}
	}
	s.Save()
}

func (s *Shapes) Save() {
	Save(s, utils.ShapesPath())
}
