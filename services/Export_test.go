package services

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/wendehals/bricks/model"
	"github.com/wendehals/bricks/test"
	"github.com/wendehals/bricks/utils"
)

func Test_ExportCollectionToHTML(t *testing.T) {
	collection := model.Load(model.NewCollection(), "test_resources/886-1.parts")
	tmpPath := filepath.FromSlash(fmt.Sprintf("%s/bricks_export", os.TempDir()))

	ExportCollectionToHTML(collection, tmpPath, "886-1")

	htmlPath := filepath.FromSlash(fmt.Sprintf("%s/886-1.html", tmpPath))

	test.AssertTrue(t, utils.FileExists(htmlPath))

	content, err := os.ReadFile(htmlPath)
	html := string(content)

	test.AssertNoError(t, err)
	test.AssertTrue(t, matches(html, "<title>Parts Collection</title>"))
	test.AssertTrue(t, matches(html, "<a href=\"https://rebrickable.com/sets/886-1/space-buggy/\">886-1 Space Buggy</a>"))
	test.AssertTrue(t, matches(html, "<p>Total number of parts: 15"))
	test.AssertTrue(t, matches(html, "title=\"Tyre 15 x 6 Offset Tread Small, Black\" alt=\"Tyre 15 x 6 Offset Tread Small\" width=\"100\" height=\"100\"/>"))

	os.RemoveAll(tmpPath)
}

func Test_ExportBuildCollectionToHTML(t *testing.T) {
	buildCollection := model.Load(model.NewBuildCollection(), "test_resources/886-1.build")
	tmpPath := filepath.FromSlash(fmt.Sprintf("%s/bricks_export", os.TempDir()))

	ExportBuildCollectionToHTML(buildCollection, tmpPath, "886-1")

	htmlPath := filepath.FromSlash(fmt.Sprintf("%s/886-1.html", tmpPath))

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

	os.RemoveAll(tmpPath)
}

func matches(text, regex string) bool {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(text)
	return match != nil
}
