package service

import (
	"github.com/paradigm-network/paradigm-fn2/api"
	"golang.org/x/net/context"
)

func Ping(ctx context.Context, msg *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{Status: "pong"}, nil
}
