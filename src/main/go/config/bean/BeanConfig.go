package bean

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
)

func NewBeanConfigInstance() *config.BeanConfig {
	return &config.BeanConfig{}
}

var beanConfigInstance = NewBeanConfigInstance()
func GetBeanConfigInstance() *config.BeanConfig {
	return beanConfigInstance
}