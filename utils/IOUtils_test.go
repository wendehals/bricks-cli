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

func Test_GetBricksDir(t *testing.T) {
	userHome, _ := os.UserHomeDir()
	expectedBricksDir := filepath.Join(userHome, ".bricks-cli")
	test.AssertSameString(t, expectedBricksDir, GetBricksDir())
}
