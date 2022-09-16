package model

import (
	"testing"

	"github.com/wendehals/bricks/test"
)

func TestLoad1(t *testing.T) {
	collection := Collection{}
	Load(&collection, "test_resources/testCollection1.parts")
	test.AssertSameInt(t, 6, len(collection.Parts))
}

func TestLoad2(t *testing.T) {
	collection := Load(&Collection{}, "test_resources/testCollection2.parts")
	test.AssertSameInt(t, 5, len(collection.Parts))
}
