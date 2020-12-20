package browserenv

import "fmt"

// shellArgs are the command and arguments needed to run a string containing
// commands in a shell.
var shellArgs = []string{"cmd", "/c"}

// shell returns a Windows `cmd` shell and "/c" argument.
func shell() (args []string) {
	return shellArgs
}

// shellEscapeCommand formats a browser command with url, ensuring url is
// properly shell-escaped.
func shellEscapeCommand(browser, url string) string {
	return fmt.Sprintf("%s %s", browser, url)
}
