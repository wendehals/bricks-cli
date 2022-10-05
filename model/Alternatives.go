package model

type Alternatives map[string][]string

func (a *Alternatives) IsCompatible(part1, part2 string) bool {
	for _, alternative := range (*a)[part1] {
		if part2 == alternative {
			return true
		}
	}

	for _, alternative := range (*a)[part2] {
		if part1 == alternative {
			return true
		}
	}

	return false
}
