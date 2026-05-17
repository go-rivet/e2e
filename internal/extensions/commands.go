package extensions

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rogpeppe/go-internal/testscript"
)

func Commands() map[string]func(ts *testscript.TestScript, neg bool, args []string) {
	return map[string]func(ts *testscript.TestScript, neg bool, args []string){
		"rivet":        rivet,
		"sleep":        sleep,
		"touch":        touch,
		"filecontains": filecontains,
	}
}

func rivet(ts *testscript.TestScript, neg bool, args []string) {
	err := ts.Exec("rivet", args...)
	ts.Check(err)
}

func filecontains(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) != 2 {
		ts.Fatalf("filecontains <file> <file>|<text>")
	}
	got := ts.ReadFile(args[0])
	want := args[1]
	if data, err := os.ReadFile(ts.MkAbs(want)); err == nil {
		want = string(data)
	}
	if strings.Contains(got, want) == neg {
		ts.Fatalf("filecontains %q; %q not found in file:\n%q", args[0], want, got)
	}
}

func sleep(ts *testscript.TestScript, neg bool, args []string) {
	duration := time.Second
	if len(args) == 1 {
		d, err := time.ParseDuration(args[0])
		ts.Check(err)
		duration = d
	}
	time.Sleep(duration)
}

func touch(ts *testscript.TestScript, neg bool, args []string) {
	if len(args) != 1 {
		ts.Fatalf("touch <file>")
	}
	// Get the relative path to the scripts current directory.
	path := ts.MkAbs(args[0])
	// Create the file (if necessary).
	err := os.MkdirAll(filepath.Dir(path), 0750)
	ts.Check(err)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	ts.Check(err)
	err = file.Close()
	ts.Check(err)
	// Now update the timestamp.
	t := time.Now()
	err = os.Chtimes(path, t, t)
	ts.Check(err)
}
