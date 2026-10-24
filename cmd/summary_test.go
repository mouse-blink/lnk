package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func (suite *CLITestSuite) TestAddRecursiveTruncatesLongSummary() {
	err := suite.runCommand("init")
	suite.Require().NoError(err)
	suite.stdout.Reset()

	// 8 files with displayLimit 5 leaves 3 hidden entries.
	testDir := filepath.Join(suite.tempDir, ".config", "many")
	suite.Require().NoError(os.MkdirAll(testDir, 0755))
	for i := 0; i < displayLimit+3; i++ {
		name := filepath.Join(testDir, fmt.Sprintf("file%d.conf", i))
		suite.Require().NoError(os.WriteFile(name, []byte("x"), 0644))
	}

	err = suite.runCommand("add", "--recursive", testDir)
	suite.Require().NoError(err)

	output := suite.stdout.String()
	suite.Contains(output, "Added 8 files recursively to lnk")
	suite.Contains(output, "... and 3 more files")
}
