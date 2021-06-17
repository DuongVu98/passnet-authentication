package app

import (
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

var appContext = context.Background()
var ctx, client, _ = okta.NewClient(
	appContext, okta.WithOrgUrl("https://dev-96211074.okta.com"),
	okta.WithToken("00wO5vZ9392pZRdjAOgHQk_D6LbI3Q1Es1fkqpm7NN"),
	okta.WithClientId("0oaw0f8g35USGAHtB5d6"),
)

func OktaClient() *okta.Client {
	return client
}

func OktaContext() context.Context {
	return ctx
}