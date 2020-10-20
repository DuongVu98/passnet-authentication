package config

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/bean"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config/handles"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
)


func RunAppConfig() {
	singletonFactory := config.GetSingletonFactory()

	//Push config to channel
	appConfigInstance := app.GetAppConfigInstance()
	singletonFactory.Set("appConfig", appConfigInstance)

	appConfigChannel := channels.GetAppConfigChannel()
	go func() {
		appConfigChannel <- appConfigInstance
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
