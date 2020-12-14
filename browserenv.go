package browserenv

import (
	"io"
	"os"

	"github.com/pkg/browser"
)

var Stderr io.Writer = os.Stderr
var Stdout io.Writer = os.Stdout

func OpenFile(path string) error {
	return browser.OpenFile(path)
}

func OpenReader(r io.Reader) error {
	return browser.OpenReader(r)
}

func OpenURL(url string) error {
	return browser.OpenURL(url)
}
