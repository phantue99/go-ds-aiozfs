//go:build !plan9

package aiozfs

import "os"

var rename = os.Rename
