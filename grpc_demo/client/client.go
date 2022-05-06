// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-06 15:02
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/majunmin/base-go/proto"
)

var (
	OpenTLS = true
)

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

func main() {
	var opts []grpc.DialOption

	if OpenTLS {
		creds, err := credentials.NewClientTLSFromFile("./grpc_demo/etc/server.pem", "localhost")
		if err != nil {
			panic(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	// 创建 TLS 连接
	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	c := proto.NewGreeterClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greeting: %s", r.Message)

	listValue, err := c.ListValue(ctx, &proto.SimpleRequest{Data: "majm"})

	if err != nil {
		log.Fatalf("could not get stream value: %v", err)
	}

	for {
		recv, err := listValue.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(recv.GetStreamValue())
	}

}
