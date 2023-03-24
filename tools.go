//go:build tools

//go:generate sh -c "GOBIN=$PWD/.gobincache go install $(sed -n 's/.*_ \"\\(.*\\)\".*/\\1/p' <$GOFILE)"
package tools

import (
	_ "ariga.io/entimport/cmd/entimport"
	_ "entgo.io/ent/cmd/ent"

	_ "github.com/cosmtrek/air"
	_ "github.com/mattn/goveralls"
	_ "golang.org/x/tools/cmd/stringer"
	_ "gotest.tools/gotestsum"
)
