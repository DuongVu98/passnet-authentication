package executor

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
)

type CommandExecutor interface {
	Execute(requestContext context.Context, command command.BaseCommand) (aggregate.User, error)
}

type CompensatingExecutor interface {
	Rollback(requestContext context.Context, command command.BaseCompensating) error
}
