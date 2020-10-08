package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func GrpcConfig(){
	grpcPort := os.Getenv("GRPC_SERVER_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	log.Printf("before grpc listen")
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()
	log.Printf("after grpc listen")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server on port %s: %v", grpcPort, err)
	}

}
