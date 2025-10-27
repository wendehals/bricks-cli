package model

import (
	"testing"

	"github.com/wendehals/bricks-cli/test"
)

func Test_Shape_Dimensions(t *testing.T) {
	shape := Shape{}

	shape.Name = "Brick 1 x 4"
	test.AssertSameString(t, "", shape.Dimensions())

	shape.Name = "Brick 1 x 8"
	test.AssertSameString(t, "8", shape.Dimensions())

	shape.Name = "Brick 2 x 8"
	test.AssertSameString(t, "2 x 8", shape.Dimensions())

	shape.Name = "Technic Brick 1 x 4"
	test.AssertSameString(t, "", shape.Dimensions())

	shape.Name = "Technic Brick 1 x 8"
	test.AssertSameString(t, "8", shape.Dimensions())

	shape.Name = "Plate 1 x 2"
	test.AssertSameString(t, "", shape.Dimensions())

	shape.Name = "Plate 1 x 6"
	test.AssertSameString(t, "6", shape.Dimensions())

	shape.Name = "Plate 4 x 6"
	test.AssertSameString(t, "4 x 6", shape.Dimensions())

	shape.Name = "Technic Plate 2 x 4"
	test.AssertSameString(t, "", shape.Dimensions())

	shape.Name = "Technic Plate 2 x 6"
	test.AssertSameString(t, "2 x 6", shape.Dimensions())

	shape.Name = "Tile 1 x 2"
	test.AssertSameString(t, "1 x 2", shape.Dimensions())

	shape.Name = "Tile 1 x 6"
	test.AssertSameString(t, "1 x 6", shape.Dimensions())

	shape.Name = "Tile 6 x 6"
	test.AssertSameString(t, "6 x 6", shape.Dimensions())

	shape.Name = "Technic Axle 2"
	test.AssertSameString(t, "2", shape.Dimensions())

	shape.Name = "Technic Axle 5.5"
	test.AssertSameString(t, "5.5", shape.Dimensions())

	shape.Name = "Technic Axle and Pin Connector Angled #2"
	test.AssertSameString(t, "#2", shape.Dimensions())

	shape.Name = "Technic Panel Fairing # 1"
	test.AssertSameString(t, "#1", shape.Dimensions())

	shape.Name = "Technic Panel Fairing #6"
	test.AssertSameString(t, "#6", shape.Dimensions())

	shape.Name = "Technic Beam 1 x 3 Thick"
	test.AssertSameString(t, "", shape.Dimensions())

	shape.Name = "Technic Beam 1 x 15 Thick"
	test.AssertSameString(t, "15", shape.Dimensions())

	shape.Name = "Technic Beam 1 x 7 Bent (4 - 4) Thick"
	test.AssertSameString(t, "", shape.Dimensions())

	shape.Name = "Hose, Pneumatic 4mm D. 33L / 26.4cm"
	test.AssertSameString(t, "26.4cm", shape.Dimensions())

	shape.Name = "Dish 2 x 2 Inverted [Radar]"
	test.AssertSameString(t, "2", shape.Dimensions())

	shape.Name = "Axle Hose, Soft 11L"
	test.AssertSameString(t, "11", shape.Dimensions())

	shape.Name = "String Cord Thin 50cm"
	test.AssertSameString(t, "50cm", shape.Dimensions())

	shape.Name = "Cone 2 x 2 x 2 with Completely Open Stud"
	test.AssertSameString(t, "", shape.Dimensions())
}
