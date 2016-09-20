package handler

import (
	"errors"

	"github.com/webx-top/db"
	"github.com/webx-top/echo"
	"github.com/webx-top/validation"

	"github.com/webx-top/echodb_api_demo/app/dbschema"
	"github.com/webx-top/echodb_api_demo/app/model"
)

func NoticeList(c echo.Context) error {
	notice := model.NewNotice(c)
	page, size := Paging(c)
	data, cnt, err := notice.ListData(nil, page, size)
	if err != nil {
		return JSON(c, nil, err)
	}
	return JSONList(c, data, cnt, page, size)
}

func NoticeAdd(c echo.Context) error {
	m := &dbschema.Notice{}
	c.Bind(m)
	valid := validation.New()
	if r := valid.Min(m.ClientId, 1, `ClientId`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	if r := valid.Min(m.EventId, 1, `EventId`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	m.Created = 0
	m.Finished = 0
	m.Retry = 0
	m.Id = 0
	_, err := m.Add()
	return JSON(c, nil, err)
}

func NoticeEdit(c echo.Context) error {
	id := c.Formx(`id`).Uint()
	if id < 1 {
		return JSON(c, nil, errors.New(`参数id错误`))
	}
	m := &dbschema.Notice{}
	c.Bind(m)
	valid := validation.New()
	if r := valid.Min(m.ClientId, 1, `ClientId`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	if r := valid.Min(m.EventId, 1, `EventId`); !r.Ok {
		return JSON(c, nil, r.Error.WithField())
	}
	m.Created = 0
	m.Finished = 0
	m.Retry = 0
	m.Id = 0
	err := m.Edit(nil, db.Cond{"id": id})
	return JSON(c, nil, err)
}

func NoticeDel(c echo.Context) error {
	id := c.Formx(`id`).Uint()
	if id < 1 {
		return JSON(c, nil, errors.New(`参数id错误`))
	}
	m := &dbschema.Notice{}
	err := m.Delete(nil, db.Cond{"id": id})
	return JSON(c, nil, err)
}
