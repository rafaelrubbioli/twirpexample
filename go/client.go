package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"twirpexample/go/proto"
)

func main() {
	ctx := context.Background()
	client := proto.NewExampleServerProtobufClient("http://localhost:6565", http.DefaultClient)

	fmt.Println("Sending message 'Hello, server!'")
	myFunctionResult, err := client.MyFunction(ctx, &proto.MyFunctionRequest{Client: proto.Client_GO, Message: "Hello, server!"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Received response from server:")
	fmt.Println(myFunctionResult.Message)

	fmt.Println("=============")
	time.Sleep(10 * time.Second)

	fmt.Println("Received response from server:")
	fmt.Println("Asking for sum (20, 30)")
	sumResult, err := client.Sum(ctx, &proto.SumRequest{First: 20, Second: 30})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response: ", sumResult.Result)
}
