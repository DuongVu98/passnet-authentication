package chain

import (
	"context"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/aggregate"
	"github.com/DuongVu98/passnet-authentication/src/main/go/domain/command"
	"github.com/satori/go.uuid"
	"log"
)

func (step PrepareEventIdStep) Execute(requestContext context.Context, c command.BaseCommand) (aggregate.User, error) {
	var newEventId = uuid.NewV4().String()
	log.Printf("generate eventId: %s", newEventId)
	requestContext = context.WithValue(requestContext, "eventId", newEventId)
	return step.Executor.Execute(requestContext, c)
}
