package server

const RpcPort = "50051"


func AddArgs(args []int) int {
	var num int
	for _, a := range args {
		num += a
	}
	return num
}


