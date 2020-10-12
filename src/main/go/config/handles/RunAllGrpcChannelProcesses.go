package handles

import grpcPackage "github.com/DuongVu98/passnet-authentication/src/main/go/adapter/grpc"

func RunAllGrpcChannelProcesses () {
	go func() {
		for {
			grpcPackage.ProcessMessage()
		}
	}()
}
