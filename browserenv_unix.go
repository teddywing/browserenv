// +build !windows

package browserenv

import (
	"fmt"
	"os"
)

// TODO
func shell() (args []string) {
	shell := os.Getenv("SHELL")

	if shell == "" {
		shell = "/bin/sh"
	}

	return []string{shell, "-c"}
}

// TODO
func escapeBrowserCommand(browser, url string) string {
	// TODO: handle %s in browser command
	return fmt.Sprintf("%s '%s'", browser, url)
}
