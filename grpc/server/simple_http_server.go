package main

import (
	"github.com/XiaoyeFang/leetcode/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"strings"
)


func muxAndHttpServer() {
	//实现拦截器 grpc本身只允许设置一个拦截器，多个拦截器会报错 "The stream server interceptor was already set and may not be reset"

	tls, err := credentials.NewServerTLSFromFile("/Users/edz/GolandProjects/leetcode/grpc/proto/server.pem", "/Users/edz/GolandProjects/leetcode/grpc/proto/server.key")
	if err != nil {
		log.Fatalf("NewServerTLSFromFile err : %s \n", err)

	}
	s := grpc.NewServer(grpc.Creds(tls), grpc.StreamInterceptor(LogStreamServerInterceptor))
	pb.RegisterGreeterServer(s, &server{})
	//if err = s.Serve(listen); err != nil {
	//	log.Fatalf("s.Serve err %s", err)
	//}
	mux := getHttpMux()
	http.ListenAndServeTLS(RpcPort,
		"/Users/edz/GolandProjects/leetcode/grpc/proto/server.pem",
		"/Users/edz/GolandProjects/leetcode/grpc/proto/server.key",
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
				log.Printf("RPC request %+v \n", request)
				s.ServeHTTP(writer, request)
			} else {
				log.Printf("request %+v \n", request)
				mux.ServeHTTP(writer, request)
			}

		}),
	)

}

func getHttpMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/morning", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("yiyayiyayo"))
	})
	return mux
}
