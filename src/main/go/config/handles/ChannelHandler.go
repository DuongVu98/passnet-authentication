package handles

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/channels"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config/bean"
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
