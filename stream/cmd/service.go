package main

import (
	pb "github.com/qinxiaogit/k8sToGo/stream/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)
type StreamService struct{}

func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}


const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main(){
	listener,err := net.Listen(Network,Address)

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}


}