package runner

import (
	"errors"
	"os"

	"github.com/go-rivet/e2e/internal/extensions"
	"github.com/rogpeppe/go-internal/testscript"
)

func Run(testDir string) error {
	if _, err := os.Stat(testDir); err != nil {
		return err
	}

	t := &cliTest{}

	testscript.RunT(t, testscript.Params{
		Dir:  testDir,
		Cmds: extensions.Commands(),
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
