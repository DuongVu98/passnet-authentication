package executor

import (
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
)

type CommandExecutor interface {
	Execute(command command.BaseCommand) (aggregate.User, error)
}

type PublishEventStep struct {
	Executor CommandExecutor
	CommandExecutor
}
