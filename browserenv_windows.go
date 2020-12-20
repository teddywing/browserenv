package browserenv

import "fmt"

var shellArgs = []string{"cmd", "/c"}

// TODO
func shell() (args []string) {
	return shellArgs
}

// TODO
func escapeBrowserCommand(browser, url string) string {
	return fmt.Sprintf("%s %s", browser, url)
}
