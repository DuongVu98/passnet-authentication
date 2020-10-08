package channels

var appBeanChannel = make(chan interface{})
func GetBeanConfigChannel() chan interface{} {
	return appBeanChannel
}

var appConfigChannel = make(chan interface{})
func GetAppConfigChannel() chan interface{} {
	return appConfigChannel
}

var sendBeanChannel = make(chan int)
func GetSendBeanChannel() chan int {
	return sendBeanChannel
}