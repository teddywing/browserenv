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
func fmtBrowserCommand(browser, url string) string {
	// TODO: handle %s in browser command
	// TODO: handle single quotes in URL
	return fmt.Sprintf("%s '%s'", browser, url)
}
