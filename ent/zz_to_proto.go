package ent

import corev1 "github.com/netzchat/server/apis/core/v1"

func (c *Channel) ToProto() *corev1.Channel {
	return &corev1.Channel{
		Id:   c.ID,
		Name: c.Name,
	}
}
