package handles

import (
	"github.com/DuongVu98/passnet-authentication/src/main/config/bean"
	"github.com/DuongVu98/passnet-authentication/src/main/domain/channels"
)

var sendBeanChannel = channels.GetSendBeanChannel()
var beanConfigChannel = channels.GetBeanConfigChannel()
func ChannelHandler() {
	go func() {
		for {
			signal := <-sendBeanChannel
			if signal == 1 {
				beanConfigChannel <- bean.GetBeanConfigInstance()
			}
		}
	}()
}
