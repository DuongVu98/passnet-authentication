package chain

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/satori/go.uuid"
)

func (step PrepareEventIdStep) Execute(requestContext context.Context, c command.BaseCommand) (aggregate.User, error) {
	var newEventId = uuid.NewV4().String()
	requestContext = context.WithValue(requestContext, "eventId", newEventId)
	return step.Execute(requestContext, c)
}
