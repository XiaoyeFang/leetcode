package main

import (
	"context"
	"github.com/XiaoyeFang/leetcode/grpc/pb"
	"google.golang.org/grpc"
	"log"
)

func swaggerClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %s \n", err)
	}
	client := pb.NewGreeterClient(conn)
	stream, err := client.SayGoodMorning(context.Background())
	if err != nil {
		log.Fatalf("SayGoodMorning err: %s \n", err)

	}
	for i := 0; i < 5; i++ {
		err = stream.SendMsg(&pb.GoodMorningReq{
			Name: int32(i),
		})
		if err != nil {
			log.Fatalf("stream.SendMsg err: %s \n", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("stream.CloseAndRecv err: %s \n", err)

	}
	log.Printf(" resp: %+v \n", resp)

}
