package main

import (
	"context"
	"github.com/XiaoyeFang/leetcode/grpc/pb"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"


func main() {
	tlsConnect()
	//testOpt(5,6)
	//swaggerClient()

}

func commonClient()  {
	//客户端负责发起调用，WithInsecure是禁用安全性的链接
	conn, err := grpc.DialContext(context.Background(), address, grpc.WithBlock(), grpc.WithInsecure())
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
	//reply, err := client.SayGoodBye(context.Background(), &pb.GoodByeReq{
	//	Name: "joy",
	//})
	//if err != nil {
	//	log.Fatalf("client.SayAdd err : %s \n", err)
	//
	//}
	//for {
	//
	//	var replay = pb.GoodByeResp{}
	//	err = reply.RecvMsg(&replay)
	//	if err != nil {
	//		if err == io.EOF {
	//			return
	//		}
	//		log.Fatalf("reply.RecvMsg err : %s \n", err)
	//
	//	}
	//	fmt.Printf("reply.Trailer().Len() %+v \n", replay.Replay)
	//}
	/*
		stream, err := streamClient.Conversations(context.Background())
		if err != nil {
			log.Fatalf("get conversations stream err: %v", err)
		}
		for n := 0; n < 5; n++ {
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
	*/

	stream, err := client.SayGoodMorning(context.Background())
	if err != nil {
		log.Fatalf("'client.SayGoodMorning err %s \n", err)
	}
	for i := 0; i < 10; i=i+2 {
		err := stream.Send(&pb.GoodMorningReq{Name: int32(i)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}

	}

	resp,err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Conversations get stream err: %v", err)
	}
	// 打印返回值
	log.Println(resp.Replay)

	//stream, err := client.SayGoodNight(context.Background())
	//if err != nil {
	//	log.Fatalf("'client.SayGoodMorning err %s \n", err)
	//}
	//for i := 0; i < 10; i = i + 2 {
	//	err := stream.Send(&pb.SayGoodNightReq{Name: fmt.Sprintf("%d", i)})
	//	if err != nil {
	//		log.Fatalf("stream request err: %v", err)
	//	}
	//	log.Println("=========", i)
	//	time.Sleep(1 * time.Second)
	//	var res = &pb.SayGoodNightResp{}
	//	err = stream.RecvMsg(res)
	//	if err == io.EOF {
	//		log.Fatalf("stream request err: %v", err)
	//
	//	}
	//	if err != nil {
	//		log.Fatalf("Conversations get stream err: %v", err)
	//	}
	//	// 打印返回值
	//	log.Println(res.Replay)
	//
	//}
	//defer stream.CloseSend()

}