package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/go-rivet/e2e/internal/runner"
	"github.com/go-rivet/e2e/internal/version"
)

type Flags struct {
	TestDir string
	Version bool
}

func parseFlags() Flags {
	var flags Flags

	pflag.StringVarP(&flags.TestDir, "dir", "d", "", "directory containing .txtar tests")
	pflag.BoolVarP(&flags.Version, "version", "v", false, "display the application version and build info")
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [DIR]\n\nFlags:\n", os.Args[0])
		pflag.PrintDefaults()
	}
	pflag.Parse()

	args := pflag.Args()
	if len(args) > 0 {
		flags.TestDir = args[0]
	}

	return flags
}

func main() {
	flags := parseFlags()

	if flags.Version {
		fmt.Println(version.GetVersionWithBuildInfo())
		os.Exit(0)
	}
	if flags.TestDir == "" {
		pflag.Usage()
		os.Exit(0)
	}

	fmt.Printf("Starting E2E Test Suite on: %s\n", flags.TestDir)
	if err := runner.Run(flags.TestDir); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error executing tests: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ All E2E test scripts completed successfully!")
	os.Exit(0)
}
