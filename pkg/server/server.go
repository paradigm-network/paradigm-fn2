package server

import (
	"context"
	"net"

	"github.com/paradigm-network/paradigm-fn2/api"
	"github.com/paradigm-network/paradigm-fn2/api/service"

	"google.golang.org/grpc"
)

type Fn2 struct {
	server *grpc.Server
	listen net.Listener
}

func NewFn2ServiceServer(uri string) *Fn2 {
	listen, err := net.Listen("tcp", uri)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	s := &Fn2{
		server: server,
		listen: listen,
	}
	api.RegisterFn2ServiceServer(server, s)
	return s
}

func (f *Fn2) Start() error {
	return f.server.Serve(f.listen)
}

func (f *Fn2) Stop() {
	if f.server == nil {
		return
	}
	f.server.Stop()
	f.server = nil
}

func (f *Fn2) Call(ctx context.Context, msg *api.CallRequest) (*api.CallResponse, error) {
	return service.Call(ctx, msg)
}

func (f *Fn2) Up(ctx context.Context, msg *api.UpRequest) (*api.UpResponse, error) {
	return service.Up(ctx, msg)
}

func (f *Fn2) Down(ctx context.Context, msg *api.DownRequest) (*api.DownResponse, error) {
	return service.Down(ctx, msg)
}

func (f *Fn2) List(ctx context.Context, msg *api.ListRequest) (*api.ListResponse, error) {
	return service.List(ctx, msg)
}

func (f *Fn2) Ping(ctx context.Context, msg *api.PingRequest) (*api.PingResponse, error) {
	return service.Ping(ctx, msg)
}
