package model

type Shapes struct {
	Shapes  map[string]Shape `json:"shapes"`
	changed bool
}

func NewShapes() *Shapes {
	return &Shapes{
		Shapes:  make(map[string]Shape, 0),
		changed: false,
	}
}

func (s *Shapes) AddShape(shape Shape) {
	s.Shapes[shape.Number] = shape
	s.changed = true
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
				s.changed = true
			}
		}
	}
}

func (s *Shapes) Save(filePath string) {
	if s.changed {
		Save(s, filePath)
		s.changed = false
	}
}
