package main

import (
	"context"
	"fmt"
	pb "github.com/XiaoyeFang/leetcode/grpc/pb"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"

	"io"
	"log"
	"net"
)

const RpcPort = ":50051"

//服务端负责实现方法
type server struct {
}

//SayAdd(context.Context, *AddRequest) (*AddReply, error)
//SayGoodBye(*GoodByeReq, Greeter_SayGoodByeServer) error
//SayGoodMorning(Greeter_SayGoodMorningServer) error
//SayGoodNight(Greeter_SayGoodNightServer) error

func (s *server) SayAdd(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {

	return &pb.AddReply{
		C: req.A + req.B,
	}, nil
}

func (s *server) SayGoodBye(req *pb.GoodByeReq, sayGoodBye pb.Greeter_SayGoodByeServer) error {
	for i := 1; i < 10; i++ {
		sayGoodBye.Send(&pb.GoodByeResp{
			Replay: fmt.Sprintf("%s:%d", req.Name, i),
		})

	}

	return nil
}

func (s *server) SayGoodMorning(stream pb.Greeter_SayGoodMorningServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GoodMorningResp{
				Replay: 200,
			})
		}
		if err != nil {
			log.Fatalf("stream.SendMsg %+v \n", err)
		}
		fmt.Println("req ",req)

	}

	return nil
}

func (s *server) SayGoodNight(stream pb.Greeter_SayGoodNightServer) error {
	var n int
	for i := 0; i < 5; i++ {
		req, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}
		}
		log.Println("server: req ", req.Name)
		n++
	}
	err := stream.SendMsg(&pb.SayGoodNightResp{
		Replay: "====",
	})
	if err != nil {
		log.Fatalf("stream.SendMsg %+v \n", err)
	}
	return nil
}

//type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

// type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
func LogStreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Printf("info: %+v \n", info.FullMethod)
	err := handler(srv, ss)
	fmt.Printf("info.Server: %+v , resp %+v \n", info.IsClientStream, srv)
	//err =ss.SendMsg(&pb.GoodMorningReq{
	//	Name: 101,
	//})
	//if err != nil {
	//	log.Fatalf("ss.RecvMsg : %s \n",err)
	//}


	return err
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//fmt.Printf("req: %+v \n", req)
	fmt.Printf("info111111111111111111111111111111: %+v \n", info.FullMethod)

	resp, err = handler(context.Background(), req)
	fmt.Printf("info.Server: %+v , resp %+v \n", info.FullMethod, resp)
	return resp, err
}
type c struct {
	hmap map[string]uint32
}

func main() {
	// 流式拦截器 and TLS
	streamLogServerInterceptorServer()
	// 提供http接口的rpc服务
	//muxAndHttpServer()
	//提供swagger接口
	//swaggerServer()

}

func streamLogServerInterceptorServer() {
	//实现拦截器 grpc本身只允许设置一个拦截器，多个拦截器会报错 "The stream server interceptor was already set and may not be reset"
	listen, err := net.Listen("tcp", RpcPort)
	if err != nil {
		log.Fatalf("net.Listen err : %s \n", err)
	}
	tls, err := credentials.NewServerTLSFromFile("/Users/edz/GolandProjects/leetcode/grpc/proto/server.pem", "/Users/edz/GolandProjects/leetcode/grpc/proto/server.key")
	if err != nil {
		log.Fatalf("NewServerTLSFromFile err : %s \n", err)

	}

	s := grpc.NewServer(grpc.Creds(tls), grpc.StreamInterceptor(LogStreamServerInterceptor))
	pb.RegisterGreeterServer(s, &server{})
	if err = s.Serve(listen); err != nil {
		log.Fatalf("s.Serve err %s", err)
	}
}



/*
interface IERC20 {
function transfer(address recopent uint256 amount) external;

}

contract MyContract{
 IERC20 usdt;
 address addr ="";
 constructor(IERC20 _usdt) public {
   usdt =_usdt;
}
function transferOut(address toAddr,uint amount) extenal{
 usdt.transfer(addr,toAddr);
}

}

*/
