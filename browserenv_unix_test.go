// +build !windows

package browserenv

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestBrowserCommand(t *testing.T) {
	envValue := "open -a Firefox"
	url := "https://duckduckgo.com"

	cmd := browserCommand(envValue, url)

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	browserCommand := fmt.Sprintf("%s '%s'", envValue, url)

	wantArgs := []string{shell, "-c", browserCommand}
	if !reflect.DeepEqual(cmd.Args, wantArgs) {
		t.Errorf("got args '%#v' want '%#v'", cmd.Args, wantArgs)
	}
}
