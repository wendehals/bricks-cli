package model

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
)

type PartRelationships struct {
	Alternatives map[string][]string `json:"alternatives"`
	Molds        map[string][]string `json:"molds"`
	Prints       map[string][]string `json:"prints"`
}

func ConvertPartRelationships(csvFile string) *PartRelationships {
	p := &PartRelationships{}
	p.Alternatives = make(map[string][]string)
	p.Molds = make(map[string][]string)
	p.Prints = make(map[string][]string)

	csvReader := csv.NewReader(getPartRelationshipsReader(csvFile))
	_, err := csvReader.Read() // omit header
	if err != nil {
		log.Fatal(err.Error())
	}
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
			p.addToAlternatives(record[1], record[2])
		case "M":
			p.addToMolds(record[1], record[2])
		case "P":
			p.addToPrints(record[1], record[2])
		}
	}

	return p
}

func ImportPartRelationships(filePath string) *PartRelationships {
	partRelationships := &PartRelationships{}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, partRelationships)
	if err != nil {
		log.Fatal(err)
	}

	return partRelationships
}

func (p *PartRelationships) IsAlternativeCompatible(part1, part2 string) bool {
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

func (p *PartRelationships) IsMoldCompatible(part1, part2 string) bool {
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

func (p *PartRelationships) IsPrintCompatible(part1, part2 string) bool {
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

func getPartRelationshipsReader(csvFile string) *gzip.Reader {
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	return gzReader
}
