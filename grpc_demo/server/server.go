// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-06 14:54
package main

import (
	"context"
	"github.com/majunmin/base-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

var (
	Address = "localhost:50051"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "hello: " + request.GetName(),
	}, nil
}

// ListValue 实现ListValue方法
func (s *server) ListValue(req *proto.SimpleRequest, srv proto.Greeter_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&proto.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		panic(err)
	}

	creds, err := credentials.NewServerTLSFromFile("./grpc_demo/etc/server.pem", "./grpc_demo/etc/server.key")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	// 注册Greeter服务
	proto.RegisterGreeterServer(s, &server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	grpclog.Println("Listen on " + Address + " with TLS")
	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
