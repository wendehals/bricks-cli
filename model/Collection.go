package model

import (
	"io"
	"log"
	"sort"
	"strconv"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks/utils"
)

const (
	CLONING_FAILED_MSG = "cloning collection failed: %s"
)

// Collection represents any collection of parts (set, parts list, ...)
type Collection struct {
	User  string   `json:"user"`
	Names []string `json:"names"`
	Sets  []Set    `json:"sets"`
	Parts []Part   `json:"parts"`
}

func NewCollection() *Collection {
	c := &Collection{}
	c.Names = []string{}
	c.Sets = []Set{}
	c.Parts = []Part{}

	return c
}

// SortByColorAndName the sets by their number and the parts their color and name.
func (c *Collection) SortByColorAndName(descending bool) *Collection {
	sort.Slice(c.Sets, func(i, j int) bool {
		return c.Sets[i].Compare(&c.Sets[j]) < 0
	})

	sort.Slice(c.Parts, func(i, j int) bool {
		if descending {
			return c.Parts[i].Compare(&c.Parts[j]) > 0
		}
		return c.Parts[i].Compare(&c.Parts[j]) < 0
	})

	return c
}

// Sort the sets by their number and the parts their quantity and name.
func (c *Collection) SortByQuantityAndName(descending bool) *Collection {
	sort.Slice(c.Sets, func(i, j int) bool {
		return c.Sets[i].Compare(&c.Sets[j]) < 0
	})

	sort.Slice(c.Parts, func(i, j int) bool {
		if descending {
			return c.Parts[i].CompareByQuantity(&c.Parts[j]) > 0
		}
		return c.Parts[i].CompareByQuantity(&c.Parts[j]) < 0
	})

	return c
}

// Find returns a PartEntry with the given part number and color id.
func (c *Collection) Find(partNum string, colorId int) *Part {
	return c.FindByEquivalence(partNum, colorId, func(s1, s2 string) bool { return s1 == s2 })
}

// FindByEquivalence returns a PartEntry with a part number that is equivalent to the given part number.
// The given function defines the equivalance.
func (c *Collection) FindByEquivalence(partNum string, colorId int, eqFunc func(string, string) bool) *Part {
	for i := range c.Parts {
		part := &c.Parts[i]
		if eqFunc(part.Shape.Number, partNum) && part.Color.ID == colorId {
			return part
		}
	}
	return nil
}

// Add one collection to another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all parts.
func (c *Collection) Add(other *Collection) *Collection {
	c.Names = []string{"Sum"}
	c.Sets = append(c.Sets, other.Sets...)

	return c.recalculateQuantity(other, utils.Add)
}

// Subtract one collection from another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all parts.
func (c *Collection) Subtract(other *Collection) *Collection {
	c.Names = []string{"Subtraction"}
	c.Sets = []Set{}

	return c.recalculateQuantity(other, utils.Subtract)
}

// Max calculates the max quantity of each part in both collections.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all parts.
func (c *Collection) Max(other *Collection) *Collection {
	c.Names = []string{"Maximum"}
	c.Sets = append(c.Sets, other.Sets...)

	return c.recalculateQuantity(other, utils.Max)
}

// CountParts sums up the quantity of all parts of a collection.
func (c *Collection) CountParts() int {
	var partsCounter int
	for _, part := range c.Parts {
		partsCounter += part.Quantity
	}
	return partsCounter
}

// HasSpareParts returns true if any part in the collection is a spare part
func (c *Collection) HasSpareParts() bool {
	for _, part := range c.Parts {
		if part.IsSpare {
			return true
		}
	}

	return false
}

// HasMissingParts returns true if any part in the collection has quantity < 0
func (c *Collection) HasMissingParts() bool {
	for _, part := range c.Parts {
		if part.Quantity < 0 {
			return true
		}
	}

	return false
}

// RemoveQuantityZero removes all parts of the collection which quantity is zero.
func (c *Collection) RemoveQuantityZero() *Collection {
	return c.Filter(func(part Part) bool { return part.Quantity != 0 })
}

// Filter applies function f on each part of the collection and removes those from the collection for which f returns false.
func (c *Collection) Filter(f func(Part) bool) *Collection {
	filteredParts := []Part{}

	for _, part := range c.Parts {
		if f(part) {
			filteredParts = append(filteredParts, part)
		}
	}

	c.Parts = filteredParts

	return c
}

// MergeByColor merges all parts of the same shape ignoring the color.
// The Color field of Part will be invalid afterwards.
// The isSpare flag of Part will be invalid afterwards and set to false for all parts.
func (c *Collection) MergeByColor() *Collection {
	c.Sets = []Set{}
	c.Names = []string{}

	partsMap := c.mapPartsByPartNumber(nil, utils.Identity)

	newParts := []Part{}
	for _, value := range partsMap {
		var currentPart = value[0]
		currentPart.Color = Color{-1, ""}

		for i := 1; i < len(value); i++ {
			currentPart.Quantity += value[i].Quantity
		}

		newParts = append(newParts, currentPart)
	}

	c.Parts = newParts

	return c
}

func (c *Collection) mapPartsByPartNumber(partsMap map[string][]Part, keyMapping func(string) string) map[string][]Part {
	if partsMap == nil {
		partsMap = make(map[string][]Part)
	}

	for _, currentPart := range c.Parts {
		currentPart.IsSpare = false
		key := keyMapping(currentPart.Shape.Number)

		value, exists := partsMap[key]
		if !exists {
			value = []Part{}
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

func (c *Collection) setParts(partsMap map[string][]Part) {
	newParts := []Part{}
	for key, value := range partsMap {
		for i := 0; i < len(value); i++ {
			value[i].Shape.Number = key
			newParts = append(newParts, value[i])
		}
	}

	c.Parts = newParts
}

func (c *Collection) recalculateQuantity(other *Collection, recalc func(int, int) int) *Collection {
	partsMap := c.mapPartsByPartNumber(nil, utils.Identity)

	for i := range other.Parts {
		found := false
		currentPart := other.Parts[i]

		values, exists := partsMap[currentPart.Shape.Number]
		if exists {
			for j := range values {
				if values[j].Color.ID == currentPart.Color.ID {
					values[j].Quantity = recalc(values[j].Quantity, currentPart.Quantity)
					found = true
					break
				}
			}
			if !found {
				currentPart.Quantity = recalc(0, currentPart.Quantity)
				currentPart.IsSpare = false
				values = append(values, currentPart)
				partsMap[currentPart.Shape.Number] = values
			}
		} else {
			currentPart.Quantity = recalc(0, currentPart.Quantity)
			currentPart.IsSpare = false
			values := []Part{currentPart}
			partsMap[currentPart.Shape.Number] = values
		}
	}

	c.setParts(partsMap)

	return c.RemoveQuantityZero()
}

// MergeAllCollections merges all given collections to a new single collection.
func MergeAllCollections(collections []*Collection) *Collection {
	log.Println("Merging parts of collections")

	mergedCollection := NewCollection()
	for _, collection := range collections {
		mergedCollection.Add(collection)
	}

	return mergedCollection
}

func ImportCSVCollection(csvFile string) *Collection {
	collection := NewCollection()

	csvReader := utils.CSVReader(csvFile)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		part := NewPart()
		part.Quantity, _ = strconv.Atoi(record[2])
		part.Shape.Number = record[0]
		part.Color.ID, _ = strconv.Atoi(record[1])

		collection.Parts = append(collection.Parts, *part)
	}

	return collection
}
