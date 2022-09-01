package model

import (
	"compress/gzip"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type PartRelationships struct {
	Alternatives map[string][]string `json:"alternatives"`
	Molds        map[string][]string `json:"molds"`
	Prints       map[string][]string `json:"prints"`
}

func ImportPartRelationships() *PartRelationships {
	p := &PartRelationships{}
	p.Alternatives = make(map[string][]string)
	p.Molds = make(map[string][]string)
	p.Prints = make(map[string][]string)

	csvReader := csv.NewReader(getPartRelationshipsReader())
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

func getPartRelationshipsReader() *gzip.Reader {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.FromSlash(fmt.Sprintf("%s/.bricks-cli/part_relationships.csv.gz", userHome))
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		DownloadPartRelationships()
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	return gzReader
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
