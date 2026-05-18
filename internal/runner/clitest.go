package runner

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/rogpeppe/go-internal/testscript"
)

type cliTest struct {
	mu      sync.Mutex
	failed  bool
	skipped bool
	isSub   bool
	parent  *cliTest
}

func (t *cliTest) Run(name string, f func(t testscript.T)) {
	subTest := &cliTest{
		isSub:  true,
		parent: t,
	}

	fmt.Printf("Running: %s\n", name)

	done := make(chan struct{})
	go func() {
		defer close(done)
		f(subTest)
	}()
	<-done

	if subTest.failed {
		t.mu.Lock()
		t.failed = true
		t.mu.Unlock()
	}
}

func (t *cliTest) Fatal(args ...interface{}) {
	t.mu.Lock()
	t.failed = true
	t.mu.Unlock()
	fmt.Printf("🚨 Fatal error: %s\n", fmt.Sprint(args...))
	if t.isSub {
		runtime.Goexit()
	}
}

func (t *cliTest) Log(args ...interface{}) {
	if t.skipped {
		fmt.Println()
	} else {
		fmt.Println(args...)
	}
}

func (t *cliTest) Logf(format string, args ...interface{}) {
	if t.skipped {
		fmt.Println()
	} else {
		fmt.Printf(format+"\n", args...)
	}
}

func (t *cliTest) Skip(args ...interface{}) {
	t.skipped = true
	fmt.Printf("⚠️  Skipped: %s\n", fmt.Sprint(args...))
	if t.isSub {
		runtime.Goexit()
	}
}

func (t *cliTest) Skipf(format string, args ...interface{}) {
	t.skipped = true
	fmt.Printf("⚠️  Skipped: %s\n", fmt.Sprintf(format, args...))
	if t.isSub {
		runtime.Goexit()
	}
}

func (t *cliTest) Errorf(format string, args ...interface{}) {
	t.mu.Lock()
	t.failed = true
	t.mu.Unlock()
	fmt.Printf("❌ "+format+"\n", args...)
}

func (t *cliTest) FailNow() {
	t.mu.Lock()
	t.failed = true
	t.mu.Unlock()
	if t.isSub {
		runtime.Goexit()
	}
}

func (t *cliTest) Parallel()     {}
func (t *cliTest) Helper()       {}
func (t *cliTest) Verbose() bool { return false }
