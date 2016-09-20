package model

import (
	"github.com/webx-top/echo"

	"github.com/webx-top/echodb_api_demo/app/dbschema"
)

func NewEvent(c echo.Context) *Event {
	return &Event{
		Event:   &dbschema.Event{},
		context: c,
	}
}

type Event struct {
	*dbschema.Event
	context echo.Context
}
