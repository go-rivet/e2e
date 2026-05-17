package runner

import (
	"fmt"
	"os"

	"github.com/rogpeppe/go-internal/testscript"
)

type cliTest struct {
	failed bool
}

func (c *cliTest) Fatal(args ...any) {
	c.failed = true
	_, _ = fmt.Fprint(os.Stderr, "❌ ")
	_, _ = fmt.Fprintln(os.Stderr, args...)
}

func (c *cliTest) Log(args ...any) {
	_, _ = fmt.Fprintln(os.Stdout, args...)
}

func (c *cliTest) FailNow() {
	c.failed = true
}

func (c *cliTest) Skip(args ...any) {
	_, _ = fmt.Fprint(os.Stdout, "⏭️ Skipped: ")
	_, _ = fmt.Fprintln(os.Stdout, args...)
}

func (c *cliTest) Parallel() {
	// No-op for CLI runner sequential runs
}

func (c *cliTest) Verbose() bool {
	return true
}

func (c *cliTest) Run(name string, f func(testscript.T)) {
	f(c)
}
