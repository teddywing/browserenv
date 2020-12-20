package browserenv

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/browser"
)

var Stderr io.Writer = os.Stderr
var Stdout io.Writer = os.Stdout

var percentS = regexp.MustCompile("%s[[:^alpha:]]?")

const commandSeparator = ":"

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
	envCommand := envBrowserCommand()
	if envCommand != "" {
		tempFile, err := ioutil.TempFile("", "browserenv")
		if err != nil {
			return err
		}

		_, err = io.Copy(tempFile, r)
		if err != nil {
			return err
		}

		return OpenFile(tempFile.Name())
	}

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
func runBrowserCommand(commands, url string) error {
	commandList := strings.Split(commands, commandSeparator)

	var err error
	for _, command := range commandList {
		cmd := browserCommand(command, url)

		// Keep running commands from left to right until one of them exits
		// successfully.
		err = cmd.Run()
		if err == nil || cmd.ProcessState.ExitCode() == 0 {
			return err
		}
	}

	return err
}

// TODO
func browserCommand(command, url string) *exec.Cmd {
	shellArgs := shell()
	shell := shellArgs[0]
	args := shellArgs[1:]

	command = fmtBrowserCommand(command, url)

	args = append(args, command)

	cmd := exec.Command(shell, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr

	return cmd
}

func fmtBrowserCommand(command, url string) string {
	url = escapeURL(url)

	if browserCommandIncludesURL(command) {
		command = fmtWithURL(command, url)
	} else {
		command = shellEscapeCommand(command, url)
	}

	return command
}

func browserCommandIncludesURL(command string) bool {
	return percentS.MatchString(command)
}

func fmtWithURL(command, url string) string {
	return strings.ReplaceAll(command, "%s", url)
}

func escapeURL(url string) string {
	return strings.ReplaceAll(url, "'", "%27")
}
