package server_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/paradigm-network/paradigm-fn2/api"
	"github.com/paradigm-network/paradigm-fn2/pkg/client"
	. "github.com/paradigm-network/paradigm-fn2/pkg/server"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var grpcEndpoint = "localhost:5001"
var cli api.Fn2ServiceClient
var server *Fn2

func startServer() {
	server = NewFn2ServiceServer(grpcEndpoint)
	go func() {
		err := server.Start()
		if err != nil {
			panic(err)
		}
	}()
	//wait for the service to start
	time.Sleep((time.Millisecond * 2000))
}

func stopServer(conn *grpc.ClientConn) {
	conn.Close()
	server.Stop()
	//wait for the service to start
	time.Sleep((time.Millisecond * 2000))
}

func TestPingService(t *testing.T) {
	ctx := context.Background()
	req := &api.PingRequest{}
	res, err := cli.Ping(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, res, &api.PingResponse{Status: "pong"})
}

func TestListService(t *testing.T) {
	ctx := context.Background()
	req := &api.ListRequest{}
	_, err := cli.List(ctx, req)
	assert.Nil(t, err)
}

func TestCallService(t *testing.T) {
	ctx := context.Background()
	req := &api.CallRequest{
		Lang: "node",
		Content: `module.exports = function (input) {
	return parseInt(input.a, 10) + parseInt(input.b, 10)
}
		`,
		Params: "a=1 b=4",
	}
	res, err := cli.Call(ctx, req)
	assert.Nil(t, err)
	assert.Equal(t, "5", res.Data)
}

func TestUpService(t *testing.T) {
	assert.Nil(t, nil)
}

func TestDownService(t *testing.T) {
	assert.Nil(t, nil)
}

func TestMain(m *testing.M) {
	startServer()

	c, conn, err := client.NewClient(grpcEndpoint)
	if err != nil {
		panic(err)
	}
	cli = c
	defer stopServer(conn)

	os.Exit(m.Run())
}
