package network

import (
	"context"
	"fmt"
	"log"
	"net"

	"user/codegen/pb/user/v1"

	"google.golang.org/grpc"
)

const port = ":30051"

type GreaterHandler struct {
	user.UnimplementedUserServer
}

func (s *GreaterHandler) SayHello(ctx context.Context, in *user.HelloRequest) (*user.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &user.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type GRPCServer struct {
	server *grpc.Server
	port   int
}

var grpc_server *GRPCServer = nil

func RunGRPCServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpc_server.port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(-1)
	}

	fmt.Println(fmt.Sprintf("GRPC PORT = :%d", grpc_server.port))

	return grpc_server.server.Serve(lis)
}

func InitGRPCServer(port int) {
	grpc_server = &GRPCServer{}
	grpc_server.port = port

	grpc_server.server = grpc.NewServer()

	user.RegisterUserServer(grpc_server.server, &GreaterHandler{})

}
