package model

import (
	"io"
	"log"
	"os"

	"github.com/wendehals/bricks/utils"
)

type PartRelationships struct {
	Alternatives map[string][]string `json:"alternatives"`
	Molds        map[string][]string `json:"molds"`
	Prints       map[string][]string `json:"prints"`
}

var (
	partRelationships *PartRelationships
)

func GetPartRelationships(update bool) *PartRelationships {
	if partRelationships != nil {
		return partRelationships
	}

	partRelationshipsPath := utils.PartRelationshipsPath()
	if utils.FileExists(partRelationshipsPath) && !update {
		partRelationships = Load(&PartRelationships{}, partRelationshipsPath)
		return partRelationships
	}

	partRelationships := &PartRelationships{}
	partRelationships.Alternatives = make(map[string][]string)
	partRelationships.Molds = make(map[string][]string)
	partRelationships.Prints = make(map[string][]string)

	csvFile := utils.DownloadPartRelationships()
	csvReader := utils.GzipCSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		switch record[0] {
		case "A":
			partRelationships.addToAlternatives(record[1], record[2])
		case "M":
			partRelationships.addToMolds(record[1], record[2])
		case "P":
			partRelationships.addToPrints(record[1], record[2])
		}
	}
	Save(partRelationships, partRelationshipsPath)
	os.Remove(csvFile)

	return partRelationships
}

func (p *PartRelationships) IsAlternative(part1, part2 string) bool {
	for _, alternative := range p.Alternatives[part1] {
		if part2 == alternative {
			return true
		}
	}

	for _, alternative := range p.Alternatives[part2] {
		if part1 == alternative {
			return true
		}
	}

	return false
}

func (p *PartRelationships) IsMold(part1, part2 string) bool {
	for _, mold := range p.Molds[part1] {
		if part2 == mold {
			return true
		}
	}

	for _, mold := range p.Molds[part2] {
		if part1 == mold {
			return true
		}
	}

	return false
}

func (p *PartRelationships) IsPrint(part1, part2 string) bool {
	for _, print := range p.Prints[part1] {
		if part2 == print {
			return true
		}
	}

	for _, print := range p.Prints[part2] {
		if part1 == print {
			return true
		}
	}

	return false
}

func (p *PartRelationships) addToAlternatives(child, parent string) {
	if p.Alternatives[child] == nil {
		p.Alternatives[child] = []string{}
	}
	p.Alternatives[child] = append(p.Alternatives[child], parent)
}

func (p *PartRelationships) addToMolds(child, parent string) {
	if p.Molds[child] == nil {
		p.Molds[child] = []string{}
	}
	p.Molds[child] = append(p.Molds[child], parent)
}

func (p *PartRelationships) addToPrints(child, parent string) {
	if p.Prints[child] == nil {
		p.Prints[child] = []string{}
	}
	p.Prints[child] = append(p.Prints[child], parent)
}
