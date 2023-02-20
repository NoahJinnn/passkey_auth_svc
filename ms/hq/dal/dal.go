// Package dal implements Data Access Layer using PostgreSQL DB.
package dal

import (
	"context"
	"fmt"
	"log"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/hq/config"
	"github.com/hellohq/hqservice/ms/hq/dal/repo"
)

const (
	schemaVersion  = 4
	dbMaxOpenConns = 100 / 10 // Use up to 1/10 of server's max_connections.
	dbMaxIdleConns = 5        // A bit more than default (2).
)

type Ctx = context.Context

type Repo struct {
	*repo.Repo
}

// New creates and returns new Repo.
// It will also run required DB migrations and connects to DB.
func New(ctx Ctx, cfg *config.PostgresConfig) (_ *Repo, err error) {

	client, err := ent.Open("postgres", "host=localhost port=5432 user=hqservice dbname=hqservice password=h3ll0HQ")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if _, err = CreateUser(ctx, client); err != nil {
		log.Fatal(err)
	}

	// TODO: Refact this part to work with entgo
	// returnErrs := []error{ // List of app.Err… returned by Repo methods.
	// 	app.ErrAlreadyExist,
	// 	app.ErrNotFound,
	// }

	r := &Repo{}
	// r.Repo, err = repo.NewPostgresRepo(ctx, repo.PostgresRepoConfig{
	// 	Postgres:   cfg,
	// 	Metric:     metric,
	// 	ReturnErrs: returnErrs,
	// })
	if err != nil {
		return nil, err
	}
	// r.DB.SetMaxOpenConns(dbMaxOpenConns)
	// r.DB.SetMaxIdleConns(dbMaxIdleConns)
	// r.SchemaVer.HoldSharedLock(ctx, time.Second)
	return r, nil
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
