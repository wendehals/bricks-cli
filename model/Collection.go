package model

import (
	"bufio"
	"embed"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/google/go-cmp/cmp"
)

// Collection represents any collection of Lego parts (set, parts list, ...)
type Collection struct {
	User  string      `json:"user"`
	IDs   []string    `json:"ids"`
	Names []string    `json:"names"`
	Parts []PartEntry `json:"parts"`
}

const (
	EXPORT_FAILED_MSG = "exporting collection to HTML file '%s' failed: %s"
)

//go:embed resources/variants.json
//go:embed resources/parts_html.gotpl
var fs embed.FS

// ImportCollection reads a collection from a JSON encoded file.
func ImportCollection(fileName string) *Collection {
	collection := &Collection{}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = json.Unmarshal(data, collection)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return collection
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
	c.IDs = append(c.IDs, other.IDs...)
	c.Names = append(c.Names, other.Names...)

	return c.recalculateQuantity(other, add)
}

// Subtract on collection from another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Subtract(other *Collection) *Collection {
	c.IDs = []string{}
	c.Names = []string{}

	c.recalculateQuantity(other, subtract)

	newParts := []PartEntry{}
	for _, partEntry := range c.Parts {
		if partEntry.Quantity != 0 {
			newParts = append(newParts, partEntry)
		}
	}
	c.Parts = newParts

	return c
}

// Max calculates the max quantity of each part in both collections.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Max(other *Collection) *Collection {
	c.IDs = []string{}
	c.Names = []string{}

	return c.recalculateQuantity(other, max)
}

// MergeByColor merges all parts of the same type ignoring the color.
// The Color field of PartEntry will be invalid afterwards.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) MergeByColor() *Collection {
	c.IDs = []string{}
	c.Names = []string{}

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
	c.IDs = []string{}
	c.Names = []string{}

	variantsMapping := loadVariants()

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

// CountParts sums up the quantity of all parts of a collection.
func (c *Collection) CountParts() int {
	var partsCounter int
	for _, partEntry := range c.Parts {
		partsCounter += partEntry.Quantity
	}
	return partsCounter
}

func (c *Collection) HasNegativePartQuantity() bool {
	for _, partEntry := range c.Parts {
		if partEntry.Quantity < 0 {
			return true
		}
	}

	return false
}

// filter applies function f on each part of the collection and removes those from the collection for which f returns false.
func (c *Collection) filter(f func(PartEntry) bool) *Collection {
	filteredParts := []PartEntry{}

	for _, partEntry := range c.Parts {
		if f(partEntry) {
			filteredParts = append(filteredParts, partEntry)
		}
	}

	c.Parts = filteredParts

	return c
}

// RemoveQuantityZero removes all parts of the collection which quantity is zero.
func (c *Collection) RemoveQuantityZero() *Collection {
	return c.filter(func(part PartEntry) bool { return part.Quantity != 0 })
}

// ExportToHTML writes an HTML file with all parts of the collection.
func (c *Collection) ExportToHTML(fileName string) {
	t := template.New("parts")
	t.Funcs(template.FuncMap{
		"abs": func(i int) int {
			if i < 0 {
				return -i
			}
			return i
		},
	})

	b, err := fs.ReadFile("resources/parts_html.gotpl")
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, fileName, err.Error())
	}

	t, err = t.Parse(string(b))
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, fileName, err.Error())
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, fileName, err.Error())
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	err = t.Execute(writer, c)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, fileName, err.Error())
	}

	writer.Flush()

	log.Printf("Exported result to '%s'\n", fileName)
}

func (c *Collection) mapPartsByPartNumber(partsMap map[string][]PartEntry, keyMapping func(string) string) map[string][]PartEntry {
	if partsMap == nil {
		partsMap = make(map[string][]PartEntry)
	}

	for _, partEntry := range c.Parts {
		partEntry.IsSpare = false

		key := keyMapping(partEntry.Part.Number)

		mappedPartEntries, exists := partsMap[key]
		if !exists {
			mappedPartEntries = []PartEntry{}
		}

		found := false
		for _, mappedPartEntry := range mappedPartEntries {
			if cmp.Equal(mappedPartEntry.Color, partEntry.Color) {
				mappedPartEntry.Quantity += partEntry.Quantity
				found = true
				break
			}
		}

		if !found {
			mappedPartEntries = append(mappedPartEntries, partEntry)
		}

		partsMap[key] = mappedPartEntries
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
	for _, currentPartEntry := range other.Parts {
		found := false

		mappedPartEntries, exists := partsMap[currentPartEntry.Part.Number]
		if exists {
			for _, mappedPartEntry := range mappedPartEntries {
				if cmp.Equal(mappedPartEntry.Color, currentPartEntry.Color) {
					mappedPartEntry.Quantity = recalc(mappedPartEntry.Quantity, currentPartEntry.Quantity)
					found = true
					break
				}
			}
		}

		if !found {
			currentPartEntry.Quantity = recalc(0, currentPartEntry.Quantity)
			currentPartEntry.IsSpare = false
			missingParts = append(missingParts, currentPartEntry)
		}
	}

	c.setParts(partsMap)
	c.Parts = append(c.Parts, missingParts...)

	return c
}

func MergeAllCollections(collections []*Collection) *Collection {
	log.Println("Merging parts of collections")

	collection := &Collection{}
	for _, collection := range collections {
		collection.Add(collection)
		collection.IDs = append(collection.IDs, collection.IDs...)
		collection.Names = append(collection.Names, collection.Names...)
	}

	return collection
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

func loadVariants() map[string]string {
	var v struct {
		Variants []struct {
			Original   string `json:"original"`
			Substitute string `json:"substitute"`
		} `json:"variants"`
	}

	jsonFile, err := fs.Open("resources/variants.json")
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

	var variantsMapping = make(map[string]string)
	for _, variant := range v.Variants {
		variantsMapping[variant.Original] = variant.Substitute
	}

	return variantsMapping
}
