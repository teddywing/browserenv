// Copyright (c) 2020  Teddy Wing
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package browserenv allows URLs and files to be opened in a local web
// browser. The system's default browser is used. If the BROWSER environment
// variable is set, the command it specifies is used instead.
//
// If the BROWSER variable contains the string "%s", that will be replaced with
// the URL. Otherwise, the URL is appended to the contents of BROWSER as its
// final argument.
//
// BROWSER can contain multiple commands delimited by colons. Each command is
// tried from left to right, stopping when a command exits with a 0 exit code.
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

// Stderr is the browser command's standard error Writer. Defaults to
// os.Stderr.
var Stderr io.Writer = os.Stderr

// Stdout is the browser command's standard output Writer. Defaults to
// os.Stdout.
var Stdout io.Writer = os.Stdout

// percentS is a regular expression that matches "%s" not followed by an
// alphabetic character.
var percentS = regexp.MustCompile("%s[[:^alpha:]]?")

// commandSeparator is the delimiter used in between multiple commands
// specified in the BROWSER environment variable.
const commandSeparator = ":"

// OpenFile opens the file referenced by path in a browser.
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

	setBrowserStdDescriptors()

	return browser.OpenFile(path)
}

// OpenReader copies the contents of r to a temporary file and opens the
// resulting file in a browser.
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

	setBrowserStdDescriptors()

	return browser.OpenReader(r)
}

// OpenURL opens url in a browser.
func OpenURL(url string) error {
	envCommand := envBrowserCommand()
	if envCommand != "" {
		return runBrowserCommand(envCommand, url)
	}

	setBrowserStdDescriptors()

	return browser.OpenURL(url)
}

// envBrowserCommand gets the value of the BROWSER environment variable.
func envBrowserCommand() string {
	return os.Getenv("BROWSER")
}

// runBrowserCommand opens url using commands, a colon-separated string of
// shell commands. Each command is executed from left to right until one exits
// with an exit code of 0.
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

// browserCommand sets up an exec.Cmd to run command, attaching Stdout and
// Stderr.
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

// fmtBrowserCommand formats command with url, producing a shell command that
// can be executed with `/bin/sh -c COMMAND`.
func fmtBrowserCommand(command, url string) string {
	url = escapeURL(url)

	if browserCommandIncludesURL(command) {
		command = fmtWithURL(command, url)
	} else {
		command = shellEscapeCommand(command, url)
	}

	return command
}

// browserCommandIncludesURL returns true if command includes a match for the
// percentS pattern.
func browserCommandIncludesURL(command string) bool {
	return percentS.MatchString(command)
}

// fmtWithURL replaces all occurrences of "%s" in command with url.
func fmtWithURL(command, url string) string {
	return strings.ReplaceAll(command, "%s", url)
}

// escapeURL replaces single quotes ("'") in url with the corresponding URL
// entity.
func escapeURL(url string) string {
	return strings.ReplaceAll(url, "'", "%27")
}

// setBrowserStdDescriptors sets browser.Stderr and browser.Stdout to Stderr
// and Stdout respectively.
func setBrowserStdDescriptors() {
	browser.Stderr = Stderr
	browser.Stdout = Stdout
}
