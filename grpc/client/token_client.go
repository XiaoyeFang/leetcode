package main

import (
	"context"
	"github.com/XiaoyeFang/leetcode/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

/*
type PerRPCCredentials interface {
	// GetRequestMetadata gets the current request metadata, refreshing
	// tokens if required. This should be called by the transport layer on
	// each request, and the data should be populated in headers or other
	// context. If a status code is returned, it will be used as the status
	// for the RPC. uri is the URI of the entry point for the request.
	// When supported by the underlying implementation, ctx can be used for
	// timeout and cancellation. Additionally, RequestInfo data will be
	// available via ctx to this call.
	// TODO(zhaoq): Define the set of the qualified keys instead of leaving
	// it as an arbitrary string.
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	// RequireTransportSecurity indicates whether the credentials requires
	// transport security.
	RequireTransportSecurity() bool
}
*/
type auth struct {
	AppKey    string
	AppSecret string
}

func (a *auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {

	return map[string]string{

	}, nil
}
func (a *auth) RequireTransportSecurity() bool {
	return true
}

func tokenClient() {

	//客户端负责发起调用，WithInsecure是禁用安全性的链接
	tls, err := credentials.NewClientTLSFromFile("/Users/edz/GolandProjects/leetcode/grpc/proto/server.pem", "localhost")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err : %s \n", err)

	}
	var auth = &auth{
		AppKey: "8hihw220f9394f98-f08uqw0ef02q3ru-23r9",
		AppSecret: "8hihw220f9394f98",
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tls),grpc.WithPerRPCCredentials(auth))
	if err != nil {
		log.Fatalf("grpc.DialContext err : %s \n", err)
	}
	client := pb.NewGreeterClient(conn)

	stream, err := client.SayGoodMorning(context.Background())
	if err != nil {
		log.Fatalf("'client.SayGoodMorning err %s \n", err)
	}
	for i := 0; i < 10; i = i + 2 {
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

/*
type Auth struct {
	appKey    string
	appSecret string
}

func (a *Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "自定义认证 Token 失败")
	}

	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return status.Errorf(codes.Unauthenticated, "自定义认证 Token 无效")
	}

	return nil
}

func (a *Auth) GetAppKey() string {
	return "eddycjy"
}

func (a *Auth) GetAppSecret() string {
	return "20181005"
}
 */