package handler

import (
	"errors"

	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/webx-top/validation"

	"github.com/webx-top/echodb_api_demo/app/dbschema"
	"github.com/webx-top/echodb_api_demo/app/model"
)

func ClientList(c echo.Context) error {
	client := model.NewClient(c)
	page, size := Paging(c)
	cnt, err := client.List(nil, nil, page, size)
	if err != nil {
		return JSON(c, nil, err)
	}
	return JSONList(c, client.Objects(), cnt, page, size)
}

func ClientAdd(c echo.Context) error {
	m := &dbschema.Client{}
	c.Bind(m)
	valid := validation.New()
	if r := valid.Required(m.Name, `Name`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	_, err := m.Add()
	return JSON(c, nil, err)
}

func ClientEdit(c echo.Context) error {
	id := c.Formx(`id`).Uint()
	if id < 1 {
		return JSON(c, nil, errors.New(`参数id错误`))
	}
	m := &dbschema.Client{}
	c.Bind(m)
	valid := validation.New()
	if r := valid.Required(m.Name, `Name`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	err := m.Edit(nil, db.Cond{"id": id})
	return JSON(c, nil, err)
}

func ClientDel(c echo.Context) error {
	id := c.Formx(`id`).Uint()
	if id < 1 {
		return JSON(c, nil, errors.New(`参数id错误`))
	}
	m := &dbschema.Client{}
	err := m.Delete(nil, db.Cond{"id": id})
	return JSON(c, nil, err)
}
