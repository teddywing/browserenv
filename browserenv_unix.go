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
func shellEscapeCommand(browser, url string) string {
	return fmt.Sprintf("%s '%s'", browser, url)
}
