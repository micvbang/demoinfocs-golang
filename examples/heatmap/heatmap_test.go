package main

import (
	"os"
	"testing"

	examples "github.com/micvbang/demoinfocs-golang/examples"
)

// Just make sure the example runs
func TestHeatmap(t *testing.T) {
	os.Args = []string{"cmd", "-demo", "../../cs-demos/default.dem"}

	examples.RedirectStdout(func() {
		main()
	})
}
