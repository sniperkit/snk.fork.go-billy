/*
Sniperkit-Bot
- Date: 2018-08-12 12:10:57.421562389 +0200 CEST m=+0.041271547
- Status: analyzed
*/

// +build windows

package test

import (
	"os"
)

var (
	customMode            os.FileMode = 0666
	expectedSymlinkTarget             = "\\dir\\file"
)
