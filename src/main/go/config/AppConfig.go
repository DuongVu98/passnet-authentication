package config

import (
	"fmt"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	grpc2 "github.com/DuongVu98/passnet-authentication/src/main/go/adapter/grpc"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func RunAppConfig() {
	go GrpcConfig()
	app.ConfigGrpcConnectionOption()
}

func GrpcConfig() {
	grpcPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()
	compensatingController := grpc2.CompensatingController{}
	myproto.RegisterCompensatingExecutorServer(grpcServer, &compensatingController)
	log.Printf("grpc listening on port %v", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server on port %s: %v", grpcPort, err)
	}
}
