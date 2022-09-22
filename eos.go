package main

import (
	"eos/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	svr := &pb.Server{}
	pb.RegisterIMMainServiceServer(s, svr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// 17a42b6596d34b9a7825d1f45a1ab97309854de85b8e982927
// 35222267ddc8438e6e57df8f3df6bbd54dbb7c4b97d755bdb1
