package config

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/bean"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/handles"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
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
	beanConfigInstance := bean.GetBeanConfigInstance()
	beanConfigInstance.UserRepository = app.GetUserMgmRepository()
	singletonFactory.Set("beanConfig", beanConfigInstance)
	// Global config
	app.SetMgmClient()

	go app.GrpcConfig()
}
