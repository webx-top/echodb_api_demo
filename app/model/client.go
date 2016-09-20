package model

import (
	"github.com/webx-top/echo"

	"github.com/webx-top/echodb_api_demo/app/dbschema"
)

func NewClient(c echo.Context) *Client {
	return &Client{
		Client:  &dbschema.Client{},
		context: c,
	}
}

type Client struct {
	*dbschema.Client
	context echo.Context
}
