package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	sdk "github.com/LoginGRPC/login-contractor"
)

const (
	address     = "http://127.0.0.1:5005/"
	defaultName = "Parmatma"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := *sdk.NewUserServiceClient(conn)

	name := defaultName

	response, err := client.Ping(context.Background(), &sdk.PingRequest{UserName: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", response.Message)
}
