package app_test

import (
	"testing"

	"github.com/powerman/check"
)

func TestRegister(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	// _, mockRepo := testNew(t)
	// var uAnon = &app.User{}
	// mockRepo.EXPECT().Create(gomock.Any(), &dom.User{})

	// tests := []struct {
	// 	user    *model.User
	// 	want    *app.User
	// 	wantErr error
	// }{
	// 	{&model.User{}, uAnon, nil},
	// 	{&model.User{}, nil, app.ErrAlreadyExist},
	// 	{&model.User{}, nil, app.ErrValidate},
	// }

	// for _, tc := range tests {
	// 	tc := tc
	// 	t.Run("", func(tt *testing.T) {
	// 		t := check.T(tt)
	// 		u, err := a.Create(ctx, tc.user)
	// 		t.Err(err, tc.wantErr)
	// 		if err == nil {
	// 			t.DeepEqual(u, tc.want)
	// 		}
	// 	})
	// }
}
