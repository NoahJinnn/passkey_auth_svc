package dal

import (
	"context"
	"fmt"
	"log"

	"github.com/hellohq/hqservice/ent"
	"github.com/hellohq/hqservice/ms/hq/app"
)

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

func (r *Repo) IncExample(ctx app.Ctx, userName string) error {
	return nil
}
