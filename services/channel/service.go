package channel

import (
	"context"

	"github.com/bufbuild/connect-go"
	corev1 "github.com/netzchat/server/apis/core/v1"
	"github.com/netzchat/server/apis/core/v1/corev1connect"
	"github.com/netzchat/server/ent"
	"github.com/netzchat/server/repositories/channel"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	corev1connect.UnimplementedChannelServiceHandler

	repository channel.Repository
}

func (s *service) List(ctx context.Context, req *connect.Request[corev1.ListChannelsRequest]) (*connect.Response[corev1.ListChannelsResponse], error) {
	channels, err := s.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&corev1.ListChannelsResponse{Channels: channels}), nil
}

func (s *service) Get(ctx context.Context, req *connect.Request[corev1.GetChannelRequest]) (*connect.Response[corev1.Channel], error) {
	channel, err := s.repository.Get(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(channel), nil
}

func (s *service) Create(ctx context.Context, req *connect.Request[corev1.CreateChannelRequest]) (*connect.Response[corev1.Channel], error) {
	channel, err := s.repository.Create(ctx, req.Msg.Channel)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(channel), nil
}

func (s *service) Update(ctx context.Context, req *connect.Request[corev1.UpdateChannelRequest]) (*connect.Response[corev1.Channel], error) {
	return nil, nil
}

func (s *service) Delete(ctx context.Context, req *connect.Request[corev1.DeleteChannelRequest]) (*connect.Response[emptypb.Empty], error) {
	err := s.repository.Delete(ctx, req.Msg.GetId())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&emptypb.Empty{}), nil
}

func New(client *ent.Client) corev1connect.ChannelServiceHandler {
	return &service{repository: channel.New(client)}
}
