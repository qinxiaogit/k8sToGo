package main

import (
	pb "github.com/qinxiaogit/k8sToGo/both/proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

var streamClient pb.StreamClient
func conversations() {
	stream, err := streamClient.Conversations(context.Background())

	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 10000; n++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}
const Address string = ":8000"

func main(){
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClient(conn)
	conversations()
}
