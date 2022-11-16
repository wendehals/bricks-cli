package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/wendehals/bricks/test"
)

func Test_CreateFolder_FileExists(t *testing.T) {
	tmpFolder := filepath.Join(os.TempDir(), "folder")
	test.AssertFalse(t, FileExists(tmpFolder))
	CreateFolder(tmpFolder)
	test.AssertTrue(t, FileExists(tmpFolder))
	os.RemoveAll(tmpFolder)
}

func Test_SplitFileName(t *testing.T) {
	base, ext := SplitFileName("test.txt")
	test.AssertSameString(t, "test", base)
	test.AssertSameString(t, ".txt", ext)

	base, ext = SplitFileName("tmp/test.txt")
	test.AssertSameString(t, "test", base)
	test.AssertSameString(t, ".txt", ext)
}

func Test_BricksDir(t *testing.T) {
	userCacheDir, _ := os.UserCacheDir()
	expectedBricksDir := filepath.Join(userCacheDir, "bricks-cli")
	test.AssertSameString(t, expectedBricksDir, CacheDir())
}

func Test_ColorsPath(t *testing.T) {
	userCacheDir, _ := os.UserCacheDir()
	expectedBricksDir := filepath.Join(userCacheDir, "bricks-cli", "colors.json")
	test.AssertSameString(t, expectedBricksDir, ColorsPath())
}

func Test_AlternatesPath(t *testing.T) {
	userCacheDir, _ := os.UserCacheDir()
	expectedBricksDir := filepath.Join(userCacheDir, "bricks-cli", "alternates.json")
	test.AssertSameString(t, expectedBricksDir, AlternatesPath())
}

func Test_MoldsPath(t *testing.T) {
	userCacheDir, _ := os.UserCacheDir()
	expectedBricksDir := filepath.Join(userCacheDir, "bricks-cli", "molds.json")
	test.AssertSameString(t, expectedBricksDir, MoldsPath())
}

func Test_PrintsPath(t *testing.T) {
	userCacheDir, _ := os.UserCacheDir()
	expectedBricksDir := filepath.Join(userCacheDir, "bricks-cli", "prints.json")
	test.AssertSameString(t, expectedBricksDir, PrintsPath())
}

func Test_ShapesPath(t *testing.T) {
	userCacheDir, _ := os.UserCacheDir()
	expectedBricksDir := filepath.Join(userCacheDir, "bricks-cli", "shapes.json")
	test.AssertSameString(t, expectedBricksDir, ShapesPath())
}

func Test_CVSReader(t *testing.T) {
	csvReader := CSVReader("test_resources/test.csv")
	record, err := csvReader.Read()
	test.AssertNoError(t, err)
	test.AssertSameString(t, "1", record[0])
	test.AssertSameString(t, "2", record[1])
	test.AssertSameString(t, "3", record[2])
}

func Test_GzipCVSReader(t *testing.T) {
	csvReader := GzipCSVReader("test_resources/test.csv.gz")
	record, err := csvReader.Read()
	test.AssertNoError(t, err)
	test.AssertSameString(t, "1", record[0])
	test.AssertSameString(t, "2", record[1])
	test.AssertSameString(t, "3", record[2])
}
