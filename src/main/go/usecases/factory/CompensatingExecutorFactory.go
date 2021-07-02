package factory

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/exception"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/executor"
	"reflect"
)

type CompensatingExecutorFactory struct {

}

func GetCompensatingExecutorFactory() CompensatingExecutorFactory {
	return CompensatingExecutorFactory{}
}

func (factory CompensatingExecutorFactory) Produce(compensating command.BaseCompensating) (executor.CompensatingExecutor, error) {
	switch reflect.TypeOf(compensating).String() {
	case reflect.TypeOf(command.RegisterCommandCompensating{}).String():
		return produce(compensating.(command.RegisterCommandCompensating)), nil
	default:
		return nil, exception.InvalidCommandException{}
	}
}

func produce(compensating command.RegisterCommandCompensating) executor.CompensatingExecutor {
	return executor.RegisterCommandExecutor{}
}
