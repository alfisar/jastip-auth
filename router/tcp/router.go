package router

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func Start() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	simpleRoute(s)
	ProfileInit().profileGrpcRouters(s)

	log.Println("gRPC server is running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
