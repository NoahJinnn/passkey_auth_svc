package provider

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent"
	"github.com/hellohq/hqservice/internal/db/sqlite/ent/connection"
)

func (p *ProviderSvc) AllConnection(ctx context.Context, userId uuid.UUID) ([]*ent.Connection, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	conns, err := entClient.Connection.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return conns, nil
}

func (p *ProviderSvc) ConnectionByProviderName(ctx context.Context, userId uuid.UUID, providerName string) (*ent.Connection, error) {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient
	conn, err := entClient.Connection.Query().Where(connection.ProviderName(providerName)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	return conn, nil
}

func (p *ProviderSvc) SaveConnection(ctx context.Context, userId uuid.UUID, providerName string, data interface{}) error {
	entClient := p.getSqliteConnect(ctx, userId.String()).entClient

	json := toJSON(data)
	_, err := entClient.Connection.Create().
		SetProviderName(providerName).
		SetData(json).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
