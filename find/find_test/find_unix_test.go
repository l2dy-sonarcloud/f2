//go:build !windows
// +build !windows

package find_test

import (
	"testing"

	"github.com/ayoisaiah/f2/v2/internal/testutil"
)

func setupWindowsHidden(_ *testing.T, _ string) (teardown func()) {
	return func() {}
}

var unixTestCases = []testutil.TestCase{
	{
		Name: "exclude hidden files by default",
		Want: []string{},
		Args: []string{"-f", "hidden", "-R"},
	},

	{
		Name: "include hidden files in search",
		Want: []string{
			".hidden_file",
			"backup/documents/.hidden_resume.txt",
			"documents/.hidden_file.txt",
			"photos/vacation/mountains/.hidden_photo.jpg",
		},
		Args: []string{"-f", "hidden", "-RH"},
	},
}

// TestFindUnix only tests search behaviors perculiar to Linux and macOS.
func TestFindUnix(t *testing.T) {
	testDir := testutil.SetupFileSystem(t, "find", findFileSystem)

	findTest(t, unixTestCases, testDir)
}
