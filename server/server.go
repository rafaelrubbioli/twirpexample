package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"twirpexample/go/proto"
)

type Server struct {}

func (s Server) MyFunction(_ context.Context, request *proto.MyFunctionRequest) (*proto.MyFunctionResponse, error) {
	return &proto.MyFunctionResponse{Message: fmt.Sprintf("Hello, %s! I recieved your message: '%s'", request.Client.Name(), request.Message)}, nil
}

func (s Server) Sum(_ context.Context, request *proto.SumRequest) (*proto.SumResponse, error) {
	return &proto.SumResponse{Result: request.First + request.Second}, nil
}

func main () {
	handler := proto.NewExampleServerServer(Server{})
	fmt.Println("Server listening on port 6565")
	if err := http.ListenAndServe(fmt.Sprintf(":6565"), handler); err != nil {
		log.Fatal(err)
	}
}
