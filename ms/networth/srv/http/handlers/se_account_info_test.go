package handlers

import (
	"context"
	"testing"

	"github.com/hellohq/hqservice/ms/networth/app/dom"
	"github.com/hellohq/hqservice/ms/networth/config"
	test "github.com/hellohq/hqservice/ms/networth/test/app"
	testRepo "github.com/hellohq/hqservice/ms/networth/test/dal"
	"github.com/stretchr/testify/assert"
)

var (
	defaultCfg = config.Config{
		SaltEdgeConfig: &config.SaltEdgeConfig{
			AppId:  "test",
			Secret: "test",
			PK:     "",
		},
	}
	ctx = context.Background()
)

func TestSeAccountInfoHandler_CreateCustomer(t *testing.T) {
	tests := []struct {
		give     *dom.CreateCustomerReq
		expected *dom.CreateCustomerResp
	}{
		{
			give: &dom.CreateCustomerReq{
				Identifier: "Josh",
			},
			expected: &dom.CreateCustomerResp{
				Identifier: "Josh",
			},
		},
	}

	repo := testRepo.NewRepo(nil)
	appl := test.NewApp(&defaultCfg, repo)
	for _, tt := range tests {
		_, err := appl.GetSeAccountInfoSvc().CreateCustomer(ctx, tt.give)
		assert.Error(t, err)
		// TODO: Need to find a way to pass down test key
		// assert.Equal(t, tt.expected.Identifier, actual.Identifier)
	}
}

func TestSeAccountInfoHandler_CreateConnectSession(t *testing.T) {

}
