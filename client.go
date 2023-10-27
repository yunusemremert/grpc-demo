package main

import (
	"context"
	"flag"
	"github.com/yunusemremert/grpc-demo/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// flag
	var name string
	flag.StringVar(&name, "name", "Yunus Emre MERT!", "Bir ad girin")

	flag.Parse()

	// conn
	var connect *grpc.ClientConn

	connect, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error when calling defer: %s", err)
		}
	}(connect)

	c := chat.NewChatServiceClient(connect)

	message := chat.HelloRequest{Name: name}

	response, responseErr := c.SayHello(context.Background(), &message)
	if responseErr != nil {
		log.Fatalf("Error when calling SayHello: %s", responseErr)
	}

	log.Printf("Response from server: %s", response.Message)
}
