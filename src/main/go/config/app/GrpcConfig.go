package app

import (
	"fmt"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func GrpcConfig() {
	grpcPort := os.Getenv("GRPC_PORT")
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

func GetSagaMessageGrpcClient() myproto.MessageServiceClient {
	var conn *grpc.ClientConn
	sagaHost := os.Getenv("GRPC_SAGA_HOST")
	sagaPort := os.Getenv("GRPC_SAGA_PORT")
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", sagaHost, sagaPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to saga server: %s", err)
	}
	return myproto.NewMessageServiceClient(conn)
}
