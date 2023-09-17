//go:build tools

package tools

import (
	_ "github.com/cosmtrek/air"
	_ "github.com/mattn/goveralls"
	_ "golang.org/x/tools/cmd/stringer"
	_ "gotest.tools/gotestsum"
)
