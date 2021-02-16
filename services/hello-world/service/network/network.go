package network

import (
	"log"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

type Network struct {
	grpc *GRPCServer
	rest *RESTServer
}

var network *Network = &Network{}

func InitNetword(rest_port int, grpc_port int) *Network {
	net := &Network{}

	InitRESTServer(rest_port)
	InitGRPCServer(grpc_port)

	net.grpc = grpc_server
	net.rest = rest_server

	return net
}

func Run() {
	g.Go(func() error {
		return RunRESTServer()
	})

	g.Go(func() error {
		return RunGRPCServer()
	})

	err := g.Wait()

	if err != nil {
		log.Fatal(err)
		panic(1)
	}
}
