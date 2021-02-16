package user

import (
	"context"
	"log"
	"time"
	"user/codegen/pb/task/v1"

	"google.golang.org/grpc"
)

type TaskService struct {
	connection *grpc.ClientConn
	client     task.TaskClient
}

func (service *TaskService) StartUp() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	service.connection = conn
	service.client = task.NewTaskClient(conn)
}

func (service *TaskService) ShutDown() {
	service.connection.Close()
}

func (service *TaskService) SayHello(name string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := service.client.SayHello(ctx, &task.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.GetMessage()
}
