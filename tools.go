//go:build tools

package tools

import (
	_ "entgo.io/ent/cmd/ent"

	_ "github.com/cosmtrek/air"
	_ "github.com/mattn/goveralls"
	_ "golang.org/x/tools/cmd/stringer"
	_ "gotest.tools/gotestsum"
)
