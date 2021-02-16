package user

import (
	"context"
	"log"
	"time"

	"user/codegen/pb/auth/v1"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type AuthService struct {
	connection *grpc.ClientConn
	client     auth.AuthClient
}

func (service *AuthService) StartUp() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	service.connection = conn
	service.client = auth.NewAuthClient(conn)
}

func (service *AuthService) ShutDown() {
	service.connection.Close()
}

func (service *AuthService) SayHello(name string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := service.client.SayHello(ctx, &auth.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.GetMessage()
}
