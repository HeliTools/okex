package api

import (
	"context"

	"github.com/HeliTools/okex"
	"github.com/HeliTools/okex/api/rest"
	"github.com/HeliTools/okex/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest       *rest.ClientRest
	Ws         *ws.ClientWs
	ctx        context.Context
	ApiKey     string
	SecretKey  string
	PassPhrase string
}

// NewClient returns a pointer to a fresh Client
func NewClient(ctx context.Context, apiKey, secretKey, passphrase string, destination okex.Destination) (*Client, error) {
	restURL := okex.RestURL
	wsPubURL := okex.PublicWsURL
	wsPriURL := okex.PrivateWsURL
	switch destination {
	case okex.AwsServer:
		restURL = okex.AwsRestURL
		wsPubURL = okex.AwsPublicWsURL
		wsPriURL = okex.AwsPrivateWsURL
	case okex.DemoServer:
		restURL = okex.DemoRestURL
		wsPubURL = okex.DemoPublicWsURL
		wsPriURL = okex.DemoPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{r, c, ctx, apiKey, secretKey, passphrase}, nil
}
