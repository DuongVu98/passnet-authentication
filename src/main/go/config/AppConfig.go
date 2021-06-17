package config

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
)

func RunAppConfig() {
	go app.GrpcConfig()
}
