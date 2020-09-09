package main

import (
	"context"
	"flag"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "p", "10001", "启动端口号")
	flag.Parse()
}

type GreeterServer struct {
}

// 服务端流式RPC
func (s *GreeterServer) SayList(req *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n < 6; n++ {
		err := stream.Send(&pb.HelloReply{
			Message: "hello.list" + strconv.Itoa(n),
		})
		if err != nil {
			return nil
		}
	}
	return nil
}

// 一元RPC
func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: req.Name,
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
