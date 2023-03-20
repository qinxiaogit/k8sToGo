package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/qinxiaogit/k8sToGo/deadlines/proto"
)

// Address 连接地址
const Address string = ":8000"

var grpcClient pb.DeadlinesClient

func main() {
	//建立连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	//建立rpc连接
	grpcClient = pb.NewDeadlinesClient(conn)
	route(ctx, 2)
}

// route 调用服务端Route方法
func route(ctx context.Context, deadlines time.Duration) {
	//设置超时
	timeDeadlines := time.Now().Add(time.Duration(deadlines * time.Second))
	ctx,cancel := context.WithDeadline(ctx,timeDeadlines)

	defer cancel()
	//创建请求结构体
	req := pb.DeadlinesRequest{
		Data:"Grpc",
	}
	// 调用我们的服务(Route方法)
	// 传入超时时间为3秒的ctx
	res,err := grpcClient.Route(ctx,&req)
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res.Value)
}
