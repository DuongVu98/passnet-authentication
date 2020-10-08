package bean

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/config"
)

func NewBeanConfigInstance() *config.BeanConfig {
	return &config.BeanConfig{}
}

var BeanConfigInstance = NewBeanConfigInstance()
func GetBeanConfigInstance() *config.BeanConfig {
	return BeanConfigInstance
}



