package main

import (
	pb "github.com/qinxiaogit/k8sToGo/both/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type StreamService struct {
}

func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,})

		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}

const (
	Address string  = ":8000"
	Network string = "tcp"
)


func main(){
	listener,err := net.Listen(Network,Address)
	if err != nil{
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServer(grpcServer,&StreamService{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}