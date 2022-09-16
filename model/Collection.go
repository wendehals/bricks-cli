package model

import (
	"archive/zip"
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks/cmd/options"
	"github.com/wendehals/bricks/utils"
)

const (
	EXPORT_FAILED_MSG  = "exporting collection to HTML file '%s' failed: %s"
	CLONING_FAILED_MSG = "cloning collection failed: %s"
)

// Collection represents any collection of parts (set, parts list, ...)
type Collection struct {
	User  string      `json:"user"`
	IDs   []string    `json:"ids"`
	Names []string    `json:"names"`
	Parts []PartEntry `json:"parts"`
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

func (c *Collection) Find(partNum string, colorId int) *PartEntry {
	return c.find(partNum, colorId, func(s1, s2 string) bool { return s1 == s2 })
}

func (c *Collection) find(partNum string, colorId int, eq func(string, string) bool) *PartEntry {
	for i := range c.Parts {
		partEntry := &c.Parts[i]
		if eq(partEntry.Part.Number, partNum) && partEntry.Color.ID == colorId {
			return partEntry
		}
	}
	return nil
}

// Add one collection to another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Add(other *Collection) *Collection {
	c.IDs = append(c.IDs, other.IDs...)
	c.Names = append(c.Names, other.Names...)

	return c.recalculateQuantity(other, add)
}

// Subtract one collection from another.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Subtract(other *Collection) *Collection {
	c.IDs = []string{}
	c.Names = []string{}

	return c.recalculateQuantity(other, subtract)
}

// Max calculates the max quantity of each part in both collections.
// The isSpare flag of PartEntry will be invalid afterwards and set to false for all entries.
func (c *Collection) Max(other *Collection) *Collection {
	c.IDs = []string{}
	c.Names = []string{}

	return c.recalculateQuantity(other, max)
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

// RemoveQuantityZero removes all parts of the collection which quantity is zero.
func (c *Collection) RemoveQuantityZero() *Collection {
	return c.filter(func(part PartEntry) bool { return part.Quantity != 0 })
}

// ExportToHTML writes an HTML file with all parts of the collection into the given export directory.
func (c *Collection) ExportToHTML(exportDir string) {
	err := os.MkdirAll(exportDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	c.replaceImageURLs(exportDir)

	htmlFileName := filepath.FromSlash(fmt.Sprintf("%s/%s.html", exportDir, exportDir))

	templ := template.New("parts")
	templ.Funcs(template.FuncMap{
		"abs": func(i int) int {
			if i < 0 {
				return -i
			}
			return i
		},
	})

	bytes, err := embeddedFS.ReadFile("resources/parts_html.gotpl")
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, htmlFileName, err.Error())
	}

	templ, err = templ.Parse(string(bytes))
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, htmlFileName, err.Error())
	}

	file, err := os.Create(htmlFileName)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, htmlFileName, err.Error())
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	err = templ.Execute(writer, c)
	if err != nil {
		log.Fatalf(EXPORT_FAILED_MSG, htmlFileName, err.Error())
	}

	writer.Flush()

	log.Printf("Exported result to '%s'\n", file.Name())
}

func (c *Collection) Build(providedParts *Collection, colors *[]Color, partRelationships *PartRelationships) *BuildMapping {
	buildMapping := NewBuildMapping()

	isMold := func(s1, s2 string) bool {
		return partRelationships.IsMold(s1, s2)
	}

	isPrint := func(s1, s2 string) bool {
		return partRelationships.IsPrint(s1, s2)
	}

	isAlternative := func(s1, s2 string) bool {
		return partRelationships.IsAlternative(s1, s2)
	}

	// search for same part type, same color
	for i := range c.Parts {
		neededPartEntry := &c.Parts[i]
		providedPartEntry := providedParts.Find(neededPartEntry.Part.Number, neededPartEntry.Color.ID)
		mapPartEntry(neededPartEntry, providedPartEntry, buildMapping)
	}

	// search for mold/print/alternative, same color
	for i := range c.Parts {
		neededPartEntry := &c.Parts[i]
		if neededPartEntry.Quantity > 0 {
			number := neededPartEntry.Part.Number
			colorId := neededPartEntry.Color.ID

			mapPartEntry(neededPartEntry, providedParts.find(number, colorId, isMold), buildMapping)
			mapPartEntry(neededPartEntry, providedParts.find(number, colorId, isPrint), buildMapping)
			mapPartEntry(neededPartEntry, providedParts.find(number, colorId, isAlternative), buildMapping)
		}
	}

	// search for same part type but different color
	for i := range c.Parts {
		neededPartEntry := &c.Parts[i]
		for j := 0; j < len(*colors) && neededPartEntry.Quantity > 0; j++ {
			colorId := (*colors)[j].ID
			if neededPartEntry.Color.ID != colorId {
				number := neededPartEntry.Part.Number
				colorId := colorId

				mapPartEntry(neededPartEntry, providedParts.Find(number, colorId), buildMapping)
				mapPartEntry(neededPartEntry, providedParts.find(number, colorId, isMold), buildMapping)
				mapPartEntry(neededPartEntry, providedParts.find(number, colorId, isPrint), buildMapping)
				mapPartEntry(neededPartEntry, providedParts.find(number, colorId, isAlternative), buildMapping)
			}
		}
	}

	// if there are neededParts with quantity > 0, then they are missing in providedParts
	for i := range c.Parts {
		neededPartEntry := &c.Parts[i]
		if neededPartEntry.Quantity > 0 {
			neededPartEntry.Quantity = -1 * neededPartEntry.Quantity
		}
	}

	return buildMapping
}

func mapPartEntry(neededPartEntry *PartEntry, providedPartEntry *PartEntry, buildMapping *BuildMapping) {
	if providedPartEntry != nil {
		mappedPartEntry := DeepClone(providedPartEntry, &PartEntry{})
		if providedPartEntry.Quantity >= neededPartEntry.Quantity {
			providedPartEntry.Quantity -= neededPartEntry.Quantity
			mappedPartEntry.Quantity = neededPartEntry.Quantity
			neededPartEntry.Quantity = 0
		} else {
			providedPartEntry.Quantity = 0
			neededPartEntry.Quantity -= mappedPartEntry.Quantity
		}
		buildMapping.Parts[*neededPartEntry] = append(buildMapping.Parts[*neededPartEntry], *mappedPartEntry)
	}
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

func (c *Collection) mapPartsByPartNumber(partsMap map[string][]PartEntry, keyMapping func(string) string) map[string][]PartEntry {
	if partsMap == nil {
		partsMap = make(map[string][]PartEntry)
	}

	for _, currentPart := range c.Parts {
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

func (c *Collection) replaceImageURLs(exportDir string) {
	for _, partEntry := range c.Parts {
		if partEntry.Color.ID >= 0 {
			imageUrl, err := extractImage(partEntry.Part.Number, partEntry.Color.ID, exportDir)
			if err != nil {
				partEntry.Part.ImageURL = imageUrl
			} else if err != nil && options.Verbose {
				log.Print(err)
			}
		}
	}
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

	return c.RemoveQuantityZero()
}

func MergeAllCollections(collections []*Collection) *Collection {
	log.Println("Merging parts of collections")

	mergedCollection := &Collection{}
	for _, collection := range collections {
		mergedCollection.Add(collection)
	}

	return mergedCollection
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

//go:embed resources/variants.json
var fs embed.FS

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

	data, err := io.ReadAll(jsonFile)
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

func extractImage(partNumber string, colorId int, exportDir string) (string, error) {
	imagesFile, err := findImagesFileForColor(colorId)
	if err != nil {
		return "", err
	}

	archive, err := zip.OpenReader(imagesFile)
	if err != nil {
		return "", err
	}
	defer archive.Close()

	imageFileName := fmt.Sprintf("%s.png", partNumber)
	for _, imageFile := range archive.File {
		if imageFile.Name == imageFileName {
			if options.Verbose {
				log.Printf("Unzipping image file '%s'", imageFile.Name)
			}

			filePath := filepath.Join(exportDir, imageFile.Name)
			dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, imageFile.Mode())
			if err != nil {
				log.Fatal(err)
			}

			fileInArchive, err := imageFile.Open()
			if err != nil {
				log.Fatal(err)
			}

			if _, err := io.Copy(dstFile, fileInArchive); err != nil {
				log.Fatal(err)
			}

			dstFile.Close()
			fileInArchive.Close()

			return imageFile.Name, nil
		}
	}

	return "", fmt.Errorf("no image file found for part number %s in color %d", partNumber, colorId)
}

func findImagesFileForColor(colorId int) (string, error) {
	imagesFile := filepath.FromSlash(fmt.Sprintf("%s/parts_%d.zip", utils.GetBricksDir(), colorId))
	if _, err := os.Stat(imagesFile); os.IsNotExist(err) {
		return "", err
	}

	return imagesFile, nil
}
