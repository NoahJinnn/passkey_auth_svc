// build integration

package dal_test

import (
	"testing"

	"github.com/powerman/check"
)

// var (
// 	errDupToken = errors.New(`duplicate key value violates unique constraint "access_tokens_pkey"`)
// 	now         = time.Now().Truncate(time.Second)
// )

func TestUser(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	// r := newTestRepo(t)
	// res, err := r.GetAllUsers(ctx)
	// t.Err(err, app.ErrNotFound)
	// t.Nil(res)

}
