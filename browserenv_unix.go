// +build !windows

package browserenv

import (
	"fmt"
	"os"
)

// shell returns the current shell specified by the SHELL environment variable
// along with a "-c" argument. If SHELL is undefined, `/bin/sh` is used.
func shell() (args []string) {
	shell := os.Getenv("SHELL")

	if shell == "" {
		shell = "/bin/sh"
	}

	return []string{shell, "-c"}
}

// shellEscapeCommand formats a browser command with url, escaping url by
// wrapping it in single quotes.
func shellEscapeCommand(browser, url string) string {
	return fmt.Sprintf("%s '%s'", browser, url)
}
