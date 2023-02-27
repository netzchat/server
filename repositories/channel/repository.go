package channel

import (
	"context"

	"github.com/netzchat/server/ent"
)

type repository struct {
	client *ent.Client
}

type Repository interface {
	Create(ctx context.Context) error
}

func New(client *ent.Client) Repository {
	return &repository{client}
}

func (r *repository) Create(ctx context.Context) error { return nil }
