package app

import (
	"flag"
	"fmt"
	myproto "github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var envErr = godotenv.Load(fmt.Sprintf("%v.env.dev", "env/"))

func GrpcConfig() {
	grpcPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcPort, err)
	}

	grpcServer := grpc.NewServer()
	log.Printf("grpc listening on port %v", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server on port %s: %v", grpcPort, err)
	}
}

var grpcClient myproto.EventProducerClient

func ConfigGrpcConnectionOption() {
	var sagaHost = os.Getenv("GRPC_SAGA_HOST")
	var sagaPort = os.Getenv("GRPC_SAGA_PORT")
	var opts []grpc.DialOption
	var serverAddr = flag.String("server_addr", fmt.Sprintf("%s:%s", sagaHost, sagaPort), "The server address in the format of host:port")

	flag.Parse()
	opts = append(opts, grpc.WithInsecure())
	//opts = append(opts, grpc.WithBlock())
	var conn, err = grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Panicf("error when Dial: %v", err)
	}
	grpcClient = myproto.NewEventProducerClient(conn)
}

func GetSagaEventProducerClient() myproto.EventProducerClient {
	return grpcClient
}
