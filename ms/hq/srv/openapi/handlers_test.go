package openapi_test

import (
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hellohq/hqservice/api/openapi/client/operations"
	"github.com/hellohq/hqservice/api/openapi/model"
	"github.com/hellohq/hqservice/ms/hq/srv/openapi"
	"github.com/powerman/check"
)

func TestHealthCheck(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	c, _, mockAppl, _ := testNewServer(t, openapi.Config{})

	mockAppl.EXPECT().HealthCheck(gomock.Any()).Return(nil, io.EOF)
	mockAppl.EXPECT().HealthCheck(gomock.Any()).Return(nil, nil)
	mockAppl.EXPECT().HealthCheck(gomock.Any()).Return("OK", nil)
	mockAppl.EXPECT().HealthCheck(gomock.Any()).Return(map[string]string{"main": "OK"}, nil)

	testCases := []struct {
		want    interface{}
		wantErr *model.Error
	}{
		{nil, apiError500},
		{nil, nil},
		{"OK", nil},
		{map[string]interface{}{"main": "OK"}, nil},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			res, err := c.Operations.HealthCheck(operations.NewHealthCheckParams())
			t.DeepEqual(openapi.ErrPayload(err), tc.wantErr)
			if res == nil {
				t.DeepEqual(nil, tc.want)
			} else {
				t.DeepEqual(res.Payload, tc.want)
			}
		})
	}
}
