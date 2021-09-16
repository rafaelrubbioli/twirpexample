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
	fmt.Printf("Recieved message: '%s'\nFrom %s\n\n", request.Message, request.Client)
	return &proto.MyFunctionResponse{Message: fmt.Sprintf("Hello, %s! I recieved your message: '%s'", request.Client.Name(), request.Message)}, nil
}

func (s Server) Sum(_ context.Context, request *proto.SumRequest) (*proto.SumResponse, error) {
	fmt.Printf("Recieved sum message: %d + %d\n", request.First, request.Second)
	return &proto.SumResponse{Result: request.First + request.Second}, nil
}

func main () {
	handler := proto.NewExampleServerServer(Server{})
	fmt.Println("Twirp server listening on port 6565")
	if err := http.ListenAndServe(fmt.Sprintf(":6565"), handler); err != nil {
		log.Fatal(err)
	}
}
