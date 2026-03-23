package app

import (
	"flag"
	"fmt"
	"io"

	"github.com/JackDrogon/go-sample/internal/version"
)

// Run executes the CLI application with the provided arguments.
func Run(args []string, stdout io.Writer) error {
	fs := flag.NewFlagSet("go-sample", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	showVersion := fs.Bool("version", false, "Print version information")
	name := fs.String("name", "go-sample", "Name to greet")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if len(fs.Args()) > 0 {
		return fmt.Errorf("unexpected arguments: %v", fs.Args())
	}

	if *showVersion {
		_, err := fmt.Fprintln(stdout, version.Info())
		return err
	}

	_, err := fmt.Fprintf(stdout, "Hello, %s!\n", *name)
	return err
}
