package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	sdk "github.com/yadavparmatma/login-contractor"
)

const (
	address     = "localhost:5005"
	defaultName = "Parmatma"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := sdk.NewGreeterClient(conn)

	name := defaultName

	response, err := client.Ping(context.Background(), &sdk.PingRequest{UserName: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.Message)
}
