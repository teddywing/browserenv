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
