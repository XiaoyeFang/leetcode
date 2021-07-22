package main

import (
	"context"
	"fmt"
	"github.com/XiaoyeFang/leetcode/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)



func tlsConnect() {

	//客户端负责发起调用，WithInsecure是禁用安全性的链接
	tls, err := credentials.NewClientTLSFromFile("/Users/edz/GolandProjects/leetcode/grpc/proto/server.pem", "localhost")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err : %s \n", err)

	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tls))
	if err != nil {
		log.Fatalf("grpc.DialContext err : %s \n", err)
	}
	client := pb.NewGreeterClient(conn)

	//resp, err := client.SayAdd(context.Background(), &pb.AddRequest{A: 3, B: 4})
	//if err != nil {
	//	log.Fatalf("client.SayAdd err : %s \n", err)
	//
	//}
	//fmt.Printf("resp %+v \n", resp.String())

	stream, err := client.SayGoodMorning(context.Background())
	if err != nil {
		log.Fatalf("'client.SayGoodMorning err %s \n", err)
	}
	for i := 1; i < 10; i = i + 2 {
		err := stream.Send(&pb.GoodMorningReq{Name: int32(i)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}

	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Conversations get stream err: %v", err)
	}
	// 打印返回值
	log.Println(resp.Replay)
}

func testOpt(arm... int)  {
	for _, a :=range arm{
		fmt.Printf(" %d \n",a)
	}
}