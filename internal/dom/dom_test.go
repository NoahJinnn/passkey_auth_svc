package dom_test

import (
	"testing"

	"github.com/hellohq/hqservice/internal/dom"
	"github.com/powerman/check"
)

func TestNewID(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()

	id1 := dom.NewID()
	id2 := dom.NewID()
	t.Match(id1, `^[a-z0-9]{26}$`)
	t.Match(id2, `^[a-z0-9]{26}$`)
	t.NotEqual(id1, id2)
}
