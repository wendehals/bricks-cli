package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wendehals/bricks/test"
)

const (
	NUMBER    = "3001"
	NAME      = "Brick 2 x 4"
	URL       = "https://rebrickable.com/parts/3001/brick-2-x-4/"
	IMAGE_URL = "https://cdn.rebrickable.com/media/thumbs/parts/ldraw/4/3001.png/250x250p.png?1658329933.4319756"
)

func Test_Add_GetShape(t *testing.T) {
	shapes := NewShapes()
	shapes.AddShape(createShape())

	actualShape, found := shapes.GetShape(NUMBER)
	test.AssertTrue(t, found)
	test.AssertSameString(t, NAME, actualShape.Name)
	test.AssertSameString(t, NUMBER, actualShape.Number)
	test.AssertSameString(t, URL, actualShape.URL)
	test.AssertSameString(t, IMAGE_URL, actualShape.ImageURL)

	_, found = shapes.GetShape("42")
	test.AssertFalse(t, found)
}

func Test_GetImageURL(t *testing.T) {
	shapes := NewShapes()
	shapes.AddShape(createShape())

	test.AssertSameString(t, IMAGE_URL, shapes.GetImageURL(NUMBER))
	test.AssertSameString(t, "", shapes.GetImageURL("42"))
}

func Test_ComplementShapesData(t *testing.T) {
	shapes := NewShapes()
	shapes.AddShape(createShape())

	shape := Shape{}
	shape.Number = "3003"
	shape.Name = "Brick 2 x 2"
	shape.URL = "https://rebrickable.com/parts/3003/brick-2-x-2/"
	shapes.Shapes[shape.Number] = shape

	parts := []Part{}

	part := NewPart()
	shape.ImageURL = "https://cdn.rebrickable.com/media/thumbs/parts/ldraw/15/3003.png/250x250p.png?1658333099.1205716"
	part.Shape = shape
	parts = append(parts, *part)

	part = NewPart()
	shape = Shape{}
	shape.Number = "3005"
	shape.Name = "Brick 1 x 1"
	shape.URL = "https://rebrickable.com/parts/3005/brick-1-x-1/"
	shape.ImageURL = "https://cdn.rebrickable.com/media/thumbs/parts/ldraw/15/3005.png/250x250p.png?1658329919.4879746"
	parts = append(parts, *part)

	shapes.ComplementShapesData(parts)
	test.AssertSameString(t, "https://cdn.rebrickable.com/media/thumbs/parts/ldraw/15/3003.png/250x250p.png?1658333099.1205716", shapes.GetImageURL("3003"))

	_, found := shapes.GetShape("3005")
	test.AssertFalse(t, found)
}

func Test_Shapes_Save(t *testing.T) {
	shapes := NewShapes()
	expectedShape := createShape()
	shapes.AddShape(expectedShape)

	tmpFilePath := filepath.Join(os.TempDir(), "shapes_tmp.json")
	defer os.Remove(tmpFilePath)

	shapes.Save(tmpFilePath)

	actualShapes, err := LoadE[Shapes](tmpFilePath)
	test.AssertNoError(t, err)

	actualShape, found := actualShapes.GetShape(expectedShape.Number)
	test.AssertTrue(t, found)
	test.AssertSameString(t, expectedShape.Number, actualShape.Number)
	test.AssertSameString(t, expectedShape.Name, actualShape.Name)
}

func createShape() Shape {
	shape := Shape{}
	shape.Number = NUMBER
	shape.Name = NAME
	shape.URL = URL
	shape.ImageURL = IMAGE_URL

	return shape
}
