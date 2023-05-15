package handlers

import (
	"context"
	"testing"

	"github.com/hellohq/hqservice/ms/networth/app/saltedge"
	"github.com/hellohq/hqservice/ms/networth/config"
	test "github.com/hellohq/hqservice/ms/networth/test/mock/app"
	testRepo "github.com/hellohq/hqservice/ms/networth/test/mock/dal"
	"github.com/stretchr/testify/assert"
)

var (
	defaultCfg = config.Config{
		SaltEdgeConfig: &config.SaltEdgeConfig{
			AppId:  "nYOGKlfdJaWf1w3rWvydX4vLjFDq8FBrhFh59yPHYJ0",
			Secret: "CN7RcowLqx6cPifqaFBEO0xeAvVn-vLf2QicECPwQNM",
			PK:     "",
		},
	}
	ctx = context.Background()
)

func TestSeAccountInfoHandler_CreateCustomer(t *testing.T) {
	tests := []struct {
		give           *saltedge.CreateCustomerReq
		expectedCreate *saltedge.CreateCustomerResp
		expectedDelete *saltedge.RemoveCustomerResp
	}{
		{
			give: &saltedge.CreateCustomerReq{
				Identifier: "Josh",
			},
			expectedCreate: &saltedge.CreateCustomerResp{
				Identifier: "Josh",
			},
			expectedDelete: &saltedge.RemoveCustomerResp{
				Deleted: true,
			},
		},
	}

	repo := testRepo.NewRepo(nil)
	appl := test.NewApp(&defaultCfg, repo)
	for _, tt := range tests {
		created, err := appl.GetSeAccountInfoSvc().CreateCustomer(ctx, tt.give)
		assert.NoError(t, err)
		// TODO: Need to find a way to pass down test private key
		assert.Equal(t, tt.expectedCreate.Identifier, created.Identifier)

		// Delete customer
		deleted, err := appl.GetSeAccountInfoSvc().RemoveCustomer(ctx, created.Id)
		assert.NoError(t, err)
		assert.Equal(t, tt.expectedDelete.Deleted, deleted.Deleted)
	}

}

func TestSeAccountInfoHandler_CreateConnectSession(t *testing.T) {

}
