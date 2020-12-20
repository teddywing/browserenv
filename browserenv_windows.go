// Copyright (c) 2020  Teddy Wing
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package browserenv

import "fmt"

// shellArgs are the command and arguments needed to run a string containing
// commands in a shell.
var shellArgs = []string{"cmd", "/c"}

// shell returns a Windows `cmd` shell and "/c" argument.
func shell() (args []string) {
	return shellArgs
}

// shellEscapeCommand formats a browser command with url, ensuring url is
// properly shell-escaped.
func shellEscapeCommand(browser, url string) string {
	return fmt.Sprintf("%s %s", browser, url)
}
