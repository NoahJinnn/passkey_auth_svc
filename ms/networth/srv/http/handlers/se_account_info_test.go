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
		SaltEdge: &config.SaltEdge{
			AppId:  "nYOGKlfdJaWf1w3rWvydX4vLjFDq8FBrhFh59yPHYJ0",
			Secret: "CN7RcowLqx6cPifqaFBEO0xeAvVn-vLf2QicECPwQNM",
			PK:     "",
		},
	}
	ctx = context.Background()
)

func TestSeAccountInfoHandler_Create_DeleteCustomer(t *testing.T) {
	tests := []struct {
		give           *saltedge.CreateCustomer
		expectedCreate *saltedge.Customer
		expectedDelete *saltedge.RemoveCustomer
	}{
		{
			give: &saltedge.CreateCustomer{
				Identifier: "Josh",
			},
			expectedCreate: &saltedge.Customer{
				Identifier: "Josh",
			},
			expectedDelete: &saltedge.RemoveCustomer{
				Deleted: true,
			},
		},
	}

	repo := testRepo.NewRepo(nil)
	appl := test.NewApp(&defaultCfg, repo)
	for _, tt := range tests {
		created, err := appl.GetSeAccountInfoSvc().CreateCustomer(ctx, tt.give)
		assert.NoError(t, err)
		assert.Equal(t, tt.expectedCreate.Identifier, created.Identifier)

		// Delete customer
		deleted, err := appl.GetSeAccountInfoSvc().RemoveCustomer(ctx, created.Id)
		assert.NoError(t, err)
		assert.Equal(t, tt.expectedDelete.Deleted, deleted.Deleted)
	}
}

func TestSeAccountInfoHandler_ShowCustomer_CreateConnectSession(t *testing.T) {
	tests := struct {
		give *saltedge.CreateConnectSession
	}{
		give: &saltedge.CreateConnectSession{
			CustomerId:           "1012221102530763642",
			IncludeFakeProviders: true,
			Consent: saltedge.Consent{
				Scopes: []string{"account_details", "transactions_details"},
			},
			Attempt: saltedge.Attempt{
				ReturnTo: "http://example.com/",
			},
		},
	}
	repo := testRepo.NewRepo(nil)
	appl := test.NewApp(&defaultCfg, repo)

	c, err := appl.GetSeAccountInfoSvc().Customer(ctx, tests.give.CustomerId)
	assert.NoError(t, err)
	assert.Equal(t, tests.give.CustomerId, c.Id)

	actual, err := appl.GetSeAccountInfoSvc().CreateConnectSession(ctx, tests.give)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual.ConnectUrl)
	assert.NotEmpty(t, actual.ExpiresAt)
}
