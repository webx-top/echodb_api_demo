package handler

import (
	"errors"

	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/webx-top/validation"

	"github.com/webx-top/echodb_api_demo/app/dbschema"
	"github.com/webx-top/echodb_api_demo/app/model"
)

func EventList(c echo.Context) error {
	event := model.NewEvent(c)
	page, size := Paging(c)
	cnt, err := event.List(nil, nil, page, size)
	if err != nil {
		return JSON(c, nil, err)
	}
	return JSONList(c, event.Objects(), cnt, page, size)
}

func EventAdd(c echo.Context) error {
	m := &dbschema.Event{}
	c.Bind(m)
	valid := validation.New()
	if r := valid.Required(m.Name, `Name`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	if r := valid.Required(m.Target, `Target`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	if r := valid.Min(m.TargetId, 1, `TargetId`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	m.Finished = 0
	m.Id = 0
	_, err := m.Add()
	return JSON(c, nil, err)
}

func EventEdit(c echo.Context) error {
	id := c.Formx(`id`).Uint()
	if id < 1 {
		return JSON(c, nil, errors.New(`参数id错误`))
	}
	m := &dbschema.Event{}
	c.Bind(m)
	valid := validation.New()
	if r := valid.Required(m.Name, `Name`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	if r := valid.Required(m.Target, `Target`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	if r := valid.Min(m.TargetId, 1, `TargetId`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	m.Finished = 0
	m.Id = 0
	err := m.Edit(nil, db.Cond{"id": id})
	return JSON(c, nil, err)
}

func EventDel(c echo.Context) error {
	id := c.Formx(`id`).Uint()
	if id < 1 {
		return JSON(c, nil, errors.New(`参数id错误`))
	}
	m := &dbschema.Event{}
	err := m.Delete(nil, db.Cond{"id": id})
	return JSON(c, nil, err)
}
