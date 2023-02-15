package app

import (
	"reflect"
	"testing"
)

const (
	TEST_ERR_MSG  = "Test error msg"
	TEST_ERR_CODE = 100
)

func TestErrorCode(t *testing.T) {
	t.Run("get msg", func(t *testing.T) {
		errCode := CodeError{
			Code: TEST_ERR_CODE,
			Msg:  TEST_ERR_MSG,
		}

		got := errCode.Error()
		if got != TEST_ERR_MSG {
			t.Errorf("got %q want %q", got, TEST_ERR_MSG)
		}
	})

	t.Run("get data", func(t *testing.T) {
		errCode := CodeError{
			Code: TEST_ERR_CODE,
			Msg:  TEST_ERR_MSG,
		}

		got := errCode.Data()
		want := &CodeError{
			Code: TEST_ERR_CODE,
			Msg:  TEST_ERR_MSG,
		}
		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
