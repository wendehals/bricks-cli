package model

type Prints struct {
	Map map[string]string `json:"prints"`
}

func NewPrints() *Prints {
	return &Prints{
		Map: map[string]string{},
	}
}

func (p *Prints) Add(part1, part2 string) {
	p.Map[part1] = part2
}

func (p *Prints) IsCompatible(part1, part2 string) bool {
	if mappedPart, found := p.Map[part1]; found && mappedPart == part2 {
		return true
	}

	if mappedPart, found := p.Map[part2]; found && mappedPart == part1 {
		return true
	}

	return false
}

func (p *Prints) RepresentativeFor(part string) string {
	if mappedPart, found := p.Map[part]; found {
		return mappedPart
	}
	return part
}
