// +build !windows

package browserenv

import (
	"os"
	"reflect"
	"testing"
)

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
