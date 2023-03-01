package main

import (
	"context"
	pb "github.com/qinxiaogit/k8sToGo/stream/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var grpcClient pb.StreamServerClient

const   Address string = ":8000"


func listValue() {
	req := pb.SimpleRequest{Data: "hello world"}

	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}
}


func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewStreamServerClient(conn)
	//route()
	listValue()
}


