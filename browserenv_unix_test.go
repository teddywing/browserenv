// +build !windows

package browserenv

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func unsetEnvBrowser(t *testing.T) {
	t.Helper()

	err := os.Unsetenv("BROWSER")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBrowserCommand(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		url      string
		wantCmd  string
	}{
		{
			"without URL directive",
			"open -a Firefox",
			"https://duckduckgo.com",
			"open -a Firefox 'https://duckduckgo.com'",
		},
		{
			"with URL directive at end",
			"open -a Firefox %s",
			"https://duckduckgo.com",
			"open -a Firefox https://duckduckgo.com",
		},
		{
			"with URL directive in middle",
			"open -a Firefox %s --other-arg",
			"https://duckduckgo.com",
			"open -a Firefox https://duckduckgo.com --other-arg",
		},
		{
			"escapes single quotes in URL",
			"open -a Firefox",
			"https://duckduckgo.com/?q='s-Hertogenbosch",
			"open -a Firefox 'https://duckduckgo.com/?q=%27s-Hertogenbosch'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := browserCommand(test.envValue, test.url)

			shell := os.Getenv("SHELL")
			if shell == "" {
				shell = "/bin/sh"
			}

			wantArgs := []string{shell, "-c", test.wantCmd}
			if !reflect.DeepEqual(cmd.Args, wantArgs) {
				t.Errorf("got args '%#v' want '%#v'", cmd.Args, wantArgs)
			}
		})
	}
}

func TestOpenURLStdout(t *testing.T) {
	var stdout strings.Builder
	Stdout = &stdout

	err := os.Setenv("BROWSER", "printf")
	if err != nil {
		t.Fatal(err)
	}

	url := "http://localhost:8000"

	OpenURL(url)

	got := stdout.String()
	if got != url {
		t.Errorf("got stdout value %q want %q", got, url)
	}

	unsetEnvBrowser(t)
}

func TestOpenURLStderr(t *testing.T) {
	var stderr strings.Builder
	Stderr = &stderr

	err := os.Setenv("BROWSER", "printf >&2")
	if err != nil {
		t.Fatal(err)
	}

	url := "http://localhost:8000"

	OpenURL(url)

	got := stderr.String()
	if got != url {
		t.Errorf("got stdout value %q want %q", got, url)
	}

	unsetEnvBrowser(t)
}

func TestOpenURLMultipleBrowserCommands(t *testing.T) {
	// The `test -z URL` command must fail, causing `printf URL` to run.
	err := os.Setenv("BROWSER", "test -z:printf")
	if err != nil {
		t.Fatal(err)
	}

	var stdout strings.Builder
	Stdout = &stdout

	url := "http://localhost:8000"

	OpenURL(url)

	got := stdout.String()
	if got != url {
		t.Errorf("got stdout value %q want %q", got, url)
	}

	unsetEnvBrowser(t)
}

func TestOpenFilePkgBrowserUsesStderr(t *testing.T) {
	var stderr strings.Builder
	Stderr = &stderr

	OpenFile("file:///tmp/does-not-exist")

	got := stderr.String()
	if got == "" {
		t.Errorf("got empty stderr want an error message")
	}

	unsetEnvBrowser(t)
}

func TestOpenURLPkgBrowserUsesStderr(t *testing.T) {
	var stderr strings.Builder
	Stderr = &stderr

	OpenURL("file:///tmp/does-not-exist")

	got := stderr.String()
	if got == "" {
		t.Errorf("got empty stderr want an error message")
	}

	unsetEnvBrowser(t)
}
