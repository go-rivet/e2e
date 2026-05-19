package runner

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-rivet/e2e/internal/extensions"
	"github.com/rogpeppe/go-internal/testscript"
)

func Run(testDir string) error {
	if _, err := os.Stat(testDir); err != nil {
		return err
	}

	var testFiles []string
	err := filepath.WalkDir(testDir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if !d.IsDir() {
			name := strings.ToLower(d.Name())
			if strings.HasSuffix(name, ".txtar") {
				testFiles = append(testFiles, path)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(testFiles) == 0 {
		return nil
	}

	t := &cliTest{}
	testscript.RunT(t, testscript.Params{
		Files: testFiles,
		Cmds:  extensions.Commands(),
		// Setup: func(e *testscript.Env) error {
		// 	var vars = []string{
		// 		fmt.Sprintf("VERSION=%s", os.Getenv("VERSION")),
		// 	}
		// 	e.Vars = append(e.Vars, vars...)
		// 	return nil
		// },
	})
	if t.failed {
		return errors.New("one or more E2E test scripts failed")
	}
	return nil
}
