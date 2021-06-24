package factory

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/chain"
	"github.com/DuongVu98/passnet-authentication/src/main/go/usecases/executor"
)

type CommandExecutorFactory struct {
}

func (cef CommandExecutorFactory) Produce(command command.RegisterCommand) executor.CommandExecutor {
	return chain.PrepareEventIdStep{
		Executor: chain.PublishEventStep{
			Executor: chain.BackupCompensatingStep{
				Executor: executor.RegisterCommandExecutor{},
			},
		},
	}
}

func GetCommandExecutorFactory() CommandExecutorFactory {
	return CommandExecutorFactory{}
}
