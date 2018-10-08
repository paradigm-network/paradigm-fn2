package service_test

import (
	"context"
	"testing"

	"github.com/paradigm-network/paradigm-fn2/api"
	. "github.com/paradigm-network/paradigm-fn2/api/service"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	ctx := context.Background()
	pingReq := &api.PingRequest{}
	pingRes, err := Ping(ctx, pingReq)
	assert.Nil(t, err)
	assert.Equal(t, pingRes, &api.PingResponse{Status: "pong"})
}
