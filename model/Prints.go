package model

type Prints struct {
	Map map[string][]string `json:"prints"`
}

func NewPrints() *Prints {
	return &Prints{
		Map: map[string][]string{},
	}
}

func (p *Prints) Add(part1, part2 string) {
	if _, found := p.Map[part1]; !found {
		p.Map[part1] = []string{}
	}
	p.Map[part1] = append(p.Map[part1], part2)
}

func (p *Prints) IsCompatible(part1, part2 string) bool {
	for _, print := range p.Map[part1] {
		if part2 == print {
			return true
		}
	}

	for _, print := range p.Map[part2] {
		if part1 == print {
			return true
		}
	}

	return false
}
