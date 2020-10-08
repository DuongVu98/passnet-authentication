package config

import (
	"github.com/DuongVu98/passnet-authentication/src/main/config/app"
	"github.com/DuongVu98/passnet-authentication/src/main/config/bean"
	"github.com/DuongVu98/passnet-authentication/src/main/config/handles"
	"github.com/DuongVu98/passnet-authentication/src/main/domain/channels"
)


func RunAppConfig() {
	//bean.GetBeanContainer()
	handles.ChannelHandler()

	//Push config to channel
	appConfigIntance := app.GetAppConfigInstance()
	appConfigChannel := channels.GetAppConfigChannel()
	go func() {
		appConfigChannel <- appConfigIntance
	}()

	//Bean config
	bean.BeanConfigInstance.UserRepository = app.GetUserMgmRepository()
	// Global config
	app.SetMgmClient()

	go app.GrpcConfig()
}
