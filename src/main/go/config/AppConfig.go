package config

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/handles"
)

func RunAppConfig() {
	handles.RunAllGrpcChannelProcesses()
	go app.GrpcConfig()
}
