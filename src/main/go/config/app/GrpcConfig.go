package app

import (
	"flag"
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/gen/src/main/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

var grpcClient proto.EventProducerClient

func ConfigGrpcConnectionOption() {
	LoadEnv()
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
	grpcClient = proto.NewEventProducerClient(conn)
}

func GetSagaEventProducerClient() proto.EventProducerClient {
	return grpcClient
}
