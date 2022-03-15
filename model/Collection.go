package model

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"text/template"

	"github.com/google/go-cmp/cmp"
)

// Collection represents any collection of Lego parts (set, parts list, ...)
type Collection struct {
	Parts []PartEntry `json:"parts"`
}

var variantsMapping = make(map[string]string)

func _init() {
	var v struct {
		Variants []struct {
			Original   string `json:"original"`
			Substitute string `json:"substitute"`
		} `json:"variants"`
	}

	jsonFile, err := os.Open("variants.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := range v.Variants {
		variantsMapping[v.Variants[i].Original] = v.Variants[i].Substitute
	}
}

// Sort the Parts of a collection by their Part Number
func (c *Collection) Sort() *Collection {
	sort.Slice(c.Parts, func(i, j int) bool {
		r := regexp.MustCompile(`(\d+)(.*)`)
		matchi := r.FindStringSubmatch(c.Parts[i].Part.Number)
		inti, _ := strconv.Atoi(matchi[1])

		matchj := r.FindStringSubmatch(c.Parts[j].Part.Number)
		intj, _ := strconv.Atoi(matchj[1])

		if inti != intj {
			return inti < intj
		}

		return matchi[2] < matchj[2]
	})

	return c
}

// Add one collection to another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Add(other *Collection) *Collection {
	return c.recalculateQuantity(other, add)
}

// Subtract on collection from another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Subtract(other *Collection) *Collection {
	c.recalculateQuantity(other, subtract)

	newParts := []PartEntry{}
	for i := range c.Parts {
		if c.Parts[i].Quantity != 0 {
			newParts = append(newParts, c.Parts[i])
		}
	}
	c.Parts = newParts

	return c
}

// Max calculates the max quantity of each part in both collections.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Max(other *Collection) *Collection {
	return c.recalculateQuantity(other, max)
}

// MergeByColor merges all parts of the same type ignoring the color.
// The Color field of PartEntry will be invalid afterwards.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) MergeByColor() *Collection {
	partsMap := c.mapPartsByPartNumber(nil, identity)

	newParts := []PartEntry{}
	for _, value := range partsMap {
		var currentPart = value[0]
		currentPart.Color = ColorType{}

		for i := 1; i < len(value); i++ {
			currentPart.Quantity += value[i].Quantity
		}

		newParts = append(newParts, currentPart)
	}

	c.Parts = newParts

	return c
}

// MergeByVariant merges all parts of compatible types (variants).
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) MergeByVariant() *Collection {
	f := func(k string) string {
		key, exists := variantsMapping[k]
		if !exists {
			return k
		}
		return key
	}
	c.setParts(c.mapPartsByPartNumber(nil, f))

	return c
}

// Filter applies function f on each part of the collection and removes those from the collection for which f returns false.
func (c *Collection) Filter(f func(PartEntry) bool) *Collection {
	filteredParts := []PartEntry{}

	for i := range c.Parts {
		if f(c.Parts[i]) {
			filteredParts = append(filteredParts, c.Parts[i])
		}
	}

	c.Parts = filteredParts

	return c
}

// RemoveQuantityZero removes all parts of the collection which quantity is zero.
func (c *Collection) RemoveQuantityZero() *Collection {
	return c.Filter(func(part PartEntry) bool { return part.Quantity != 0 })
}

// Import reads a collection from a JSON encoded file.
func Import(fileName string) (*Collection, error) {
	var collection Collection

	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

// Export writes a collection into a JSON encoded file.
func (c *Collection) Export(fileName string) {
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Printf("Serializing to JSON failed: %s\n", err)
	}

	ioutil.WriteFile(fileName, data, os.ModePerm)
	if err != nil {
		log.Printf("Exporting collection to JSON file '%s' failed: %s\n", fileName, err)
	}

	log.Printf("Exported result to '%s'\n", fileName)
}

// ExportToHTML writes an HTML file with all parts of the collection.
func (c *Collection) ExportToHTML(fileName string) {
	t, err := template.ParseFiles("parts_html.gotpl")
	if err != nil {
		log.Printf("Exporting collection to HTML file '%s' failed: %s\n", fileName, err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Exporting collection to HTML file '%s' failed: %s\n", fileName, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	err = t.Execute(writer, c.Parts)
	if err != nil {
		log.Printf("Exporting HTML to '%s' failed: %s\n", fileName, err)
	}

	writer.Flush()

	log.Printf("Exported result to '%s'\n", fileName)
}

func (c *Collection) mapPartsByPartNumber(partsMap map[string][]PartEntry, keyMapping func(string) string) map[string][]PartEntry {
	if partsMap == nil {
		partsMap = make(map[string][]PartEntry)
	}

	for i := range c.Parts {
		currentPart := c.Parts[i]
		currentPart.IsSpare = false

		key := keyMapping(currentPart.Part.Number)

		value, exists := partsMap[key]
		if !exists {
			value = []PartEntry{}
		}

		found := false
		for j := range value {
			if cmp.Equal(value[j].Color, currentPart.Color) {
				value[j].Quantity += currentPart.Quantity
				found = true
				break
			}
		}

		if !found {
			value = append(value, currentPart)
		}

		partsMap[key] = value
	}

	return partsMap
}

func (c *Collection) setParts(partsMap map[string][]PartEntry) {
	newParts := []PartEntry{}
	for key, value := range partsMap {
		for i := 0; i < len(value); i++ {
			value[i].Part.Number = key
			newParts = append(newParts, value[i])
		}
	}

	c.Parts = newParts
}

func (c *Collection) recalculateQuantity(other *Collection, recalc func(int, int) int) *Collection {
	partsMap := c.mapPartsByPartNumber(nil, identity)

	missingParts := []PartEntry{}
	for i := range other.Parts {
		found := false
		currentPart := other.Parts[i]

		values, exists := partsMap[currentPart.Part.Number]
		if exists {
			for j := range values {
				if cmp.Equal(values[j].Color, currentPart.Color) {
					values[j].Quantity = recalc(values[j].Quantity, currentPart.Quantity)
					found = true
					break
				}
			}
		}

		if !found {
			currentPart.Quantity = recalc(0, currentPart.Quantity)
			currentPart.IsSpare = false
			missingParts = append(missingParts, currentPart)
		}
	}

	c.setParts(partsMap)
	c.Parts = append(c.Parts, missingParts...)

	return c
}

func identity(k string) string {
	return k
}

func add(x1, x2 int) int {
	return x1 + x2
}

func subtract(x1, x2 int) int {
	return x1 - x2
}

func max(x1, x2 int) int {
	if x1 >= x2 {
		return x1
	}
	return x2
}
