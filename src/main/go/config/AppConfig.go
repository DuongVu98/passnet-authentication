package config

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/handles"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config/bean"
)


func RunAppConfig() {

	//Push config to channel
	appConfigIntance := app.GetAppConfigInstance()
	singletonFactory := config.GetSingletonFactory()
	singletonFactory.Set("appConfig", appConfigIntance)

	appConfigChannel := channels.GetAppConfigChannel()
	go func() {
		appConfigChannel <- appConfigIntance
	}()

	handles.ChannelHandler()
	handles.RunAllGrpcChannelProcesses()

	//Bean config
	bean.BeanConfigInstance.UserRepository = app.GetUserMgmRepository()
	// Global config
	app.SetMgmClient()

	go app.GrpcConfig()
}
