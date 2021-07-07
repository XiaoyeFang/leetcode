package client

import (
	"google.golang.org/grpc"
)

const address = "localhost:50051"
func main() {
	conn,err :=grpc.Dial(address,grpc.WithInsecure())
}
