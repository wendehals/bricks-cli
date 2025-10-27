package services

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-cli/test"
	"github.com/wendehals/bricks-cli/utils"
)

func Test_ExportCollectionToHTML(t *testing.T) {
	collection := model.Load[model.Collection]("test_resources/886-1.parts")

	tmpPath := filepath.Join(os.TempDir(), "bricks_export")
	defer os.RemoveAll(tmpPath)

	ExportCollectionToHTML(&collection, tmpPath, "886-1")

	htmlPath := filepath.Join(tmpPath, "886-1.html")

	test.AssertTrue(t, utils.FileExists(htmlPath))

	content, err := os.ReadFile(htmlPath)
	html := string(content)

	test.AssertNoError(t, err)
	test.AssertTrue(t, matches(html, "<title>Parts Collection</title>"))
	test.AssertTrue(t, matches(html, "<a href=\"https://rebrickable.com/sets/886-1/space-buggy/\">886-1 Space Buggy</a>"))
	test.AssertTrue(t, matches(html, "<p>Total number of parts: 15"))
	test.AssertTrue(t, matches(html, "title=\"Tyre 15 x 6 Offset Tread Small, Black\" alt=\"Tyre 15 x 6 Offset Tread Small\" width=\"100\" height=\"100\"/>"))
}

func Test_ExportBuildCollectionToHTML(t *testing.T) {
	buildCollection := model.Load[model.BuildCollection]("test_resources/886-1.build")

	tmpPath := filepath.Join(os.TempDir(), "bricks_export")
	defer os.RemoveAll(tmpPath)

	ExportBuildCollectionToHTML(&buildCollection, tmpPath, "886-1")

	htmlPath := filepath.Join(tmpPath, "886-1.html")

	test.AssertTrue(t, utils.FileExists(htmlPath))

	content, err := os.ReadFile(htmlPath)
	html := string(content)

	test.AssertNoError(t, err)
	test.AssertTrue(t, matches(html, "<title>Building 886-1 Space Buggy</title>"))
	test.AssertTrue(t, matches(html, "<h1>Building <a href=\"https://rebrickable.com/sets/886-1/space-buggy/\">886-1 Space Buggy</a></h1>"))
	test.AssertTrue(t, matches(html, "<p>Total number of parts: 20"))
	test.AssertTrue(t, matches(html, "<p>Provided parts: 12"))
	test.AssertTrue(t, matches(html, "<p>Missing parts: 3"))
	test.AssertTrue(t, matches(html, "title=\"Tyre 15 x 6 Offset Tread Small, Black\" alt=\"Tyre 15 x 6 Offset Tread Small\" width=\"100\" height=\"100\"/>"))
	test.AssertTrue(t, matches(html, "<div class=\"missing\">2</div>"))
}

func matches(text, regex string) bool {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(text)
	return match != nil
}
