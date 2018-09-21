package server

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/paradigm-network/paradigm-fn2/api"
)

//Run start the HTTP server proxying request to gRPC
func Run(grpcEndpoint, listen string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	//TODO review options
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterFn2ServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(listen, mux)
}
