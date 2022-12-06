package model

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wendehals/bricks/test"
)

func Test_Load(t *testing.T) {
	collection := Load[Collection]("test_resources/testCollection1.parts")
	test.AssertSameInt(t, 6, len(collection.Parts))
}

func Test_LoadE(t *testing.T) {
	collection, err := LoadE[Collection]("test_resources/testCollection2.parts")

	test.AssertTrue(t, err == nil)
	test.AssertSameInt(t, 5, len(collection.Parts))
}

func Test_Save(t *testing.T) {
	original := Load[Collection]("test_resources/testCollection1.parts")

	tmpFile := filepath.FromSlash(fmt.Sprintf("%s/%s", os.TempDir(), "test.parts"))

	Save(original, tmpFile)

	saved := Load[Collection](tmpFile)
	os.Remove(tmpFile)

	test.AssertTrue(t, cmp.Equal(original, saved))
}

func Test_DeepClone(t *testing.T) {
	original := Load[Collection]("test_resources/testCollection1.parts")
	clone := DeepClone(original)

	test.AssertTrue(t, cmp.Equal(original, clone))
}
