browserenv
==========

[![GoDoc](https://godoc.org/github.com/teddywing/browserenv?status.svg)](https://godoc.org/github.com/teddywing/browserenv)

Browserenv allows URLs and files to be opened in a local web browser. It is a
drop-in replacement for the [`github.com/pkg/browser`][github.com/pkg/browser]
package.

If the `BROWSER` environment variable is set, the input URL will be opened using
the command it defines rather than the system's default web browser. When
`BROWSER` is not defined, `github.com/pkg/browser` is used.


## Examples
Set `BROWSER` to a command that opens a URL. The URL is appended as an argument
to the command:

	BROWSER="open -a Firefox"

If `%s` is included in the command, it is replaced with the URL:

	BROWSER="open -a Firefox '%s'"

Multiple commands can be specified, delimited by colons. The commands will be
tried from left to right, stopping when a command exits with a 0 exit code.

	BROWSER="w3m '%s':open -a Firefox"

A sample program:

``` go
package main

import (
	"strings"

	"github.com/teddywing/browserenv"
)

func main() {
	browserenv.OpenFile("file.gif")

	browserenv.OpenReader(strings.NewReader("Reader content"))

	browserenv.OpenURL("https://duckduckgo.com")
}
```


## License
Copyright © 2020 Teddy Wing. Licensed under the Mozilla Public License v. 2.0
(see the included LICENSE file).


[github.com/pkg/browser]: https://github.com/pkg/browser
