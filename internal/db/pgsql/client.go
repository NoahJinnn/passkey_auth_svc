package pgsql

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/hellohq/hqservice/ent"
)

func NewPgClient(ctx context.Context, dateSourceName string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, dateSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client
}
