package main

import (
	"context"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
)

var (
	port = "10001"
)

func main() {
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	err = SayHello(client)
	if err != nil {
		log.Fatalf("SayHello err: %s", err.Error())
	}
}

func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "test",
	})
	if err != nil {
		return err
	}
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
