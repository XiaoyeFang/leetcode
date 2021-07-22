package main

import (
	"github.com/XiaoyeFang/leetcode/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func swaggerServer() {
	listen, err := net.Listen("tcp", RpcPort)
	if err != nil {
		log.Fatalf("net.Listen err : %s \n", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("s.Serve(listen): %s \n", err)
	}

}
