package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"google.golang.org/grpc/credentials"
	"context"
	pb "github.com/qinxiaogit/k8sToGo/security/proto"
)

type SecurityService struct {

}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)


func main(){
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/tls/server.pem", "../pkg/tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryClientInterceptor

	interceptor = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//拦截普通方法请求，验证Token
		err =

		// 继续处理请求
		return handler(ctx, req)
	}
}


func Route(ctx context.Context,req *pb.SecurityRequest) (*pb.SecurityResponse, error){




}