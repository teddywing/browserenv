package browserenv

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/browser"
)

var Stderr io.Writer = os.Stderr
var Stdout io.Writer = os.Stdout

func OpenFile(path string) error {
	envCommand := envBrowserCommand()
	if envCommand != "" {
		path, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		url := "file://" + path

		return runBrowserCommand(envCommand, url)
	}

	return browser.OpenFile(path)
}

func OpenReader(r io.Reader) error {
	return browser.OpenReader(r)
}

func OpenURL(url string) error {
	envCommand := envBrowserCommand()
	if envCommand != "" {
		return runBrowserCommand(envCommand, url)
	}

	return browser.OpenURL(url)
}

// TODO
func envBrowserCommand() string {
	return os.Getenv("BROWSER")
}

// TODO
func runBrowserCommand(command, url string) error {
	return browserCommand(command, url).Run()
}

// TODO
func browserCommand(command, url string) *exec.Cmd {
	shellArgs := shell()
	shell := shellArgs[0]
	args := shellArgs[1:]

	command = fmtBrowserCommand(command, url)
	args = append(args, command)

	return exec.Command(shell, args...)
}
