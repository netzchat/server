package channel

import (
	"context"
	"fmt"

	corev1 "github.com/netzchat/server/apis/core/v1"
	"github.com/netzchat/server/ent"
	"github.com/netzchat/server/utils"
)

type repository struct {
	client *ent.Client
}

type Repository interface {
	List(context.Context) ([]*corev1.Channel, error)
	Get(context.Context, string) (*corev1.Channel, error)
	Create(context.Context, *corev1.Channel) (*corev1.Channel, error)
	// todo: Update
	Delete(context.Context, string) error
}

func New(client *ent.Client) Repository {
	return &repository{client}
}

func (r *repository) List(ctx context.Context) ([]*corev1.Channel, error) {
	channels, err := r.client.Channel.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("query channels: %w", err)
	}

	return utils.Map(func(c *ent.Channel) *corev1.Channel { return c.ToProto() }, channels), nil
}

func (r *repository) Get(ctx context.Context, id string) (*corev1.Channel, error) {
	channel, err := r.client.Channel.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get channel: %w", err)
	}

	return channel.ToProto(), nil
}

func (r *repository) Create(ctx context.Context, channel *corev1.Channel) (*corev1.Channel, error) {
	create := r.client.Channel.Create()
	create.SetName(channel.GetName())

	saved, err := create.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create channel: %w", err)
	}

	return saved.ToProto(), nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	err := r.client.Channel.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("delete channel: %w", err)
	}

	return nil
}
