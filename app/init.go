package app

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/admpub/log"
	"github.com/webx-top/echo"

	. "github.com/webx-top/echodb_api_demo/app/handler"
	"github.com/webx-top/echodb_api_demo/app/library"
	"github.com/webx-top/echodb_api_demo/app/model"
)

func Initialize(e *echo.Echo) {
	addRouter(e)
	model.MustOK(model.ParseConfig())
	me := library.MonitorEvent{
		Modify: func(file string) {
			if strings.HasSuffix(file, `.yaml`) {
				log.Info(`reload config from ` + file)
				model.MustOK(model.ParseConfig())
			}
		},
	}
	me.Watch(filepath.Dir(*model.CmdConfig.Conf))

	crontabCompleted := true
	go func() {
		for {
			if crontabCompleted {
				crontabCompleted = false
				log.Info(`[Crontab] Starting`)
				crontabCompleted = Crontab()
			}
			if model.Config.Cron.Interval <= 0 {
				model.Config.Cron.Interval = 1
			}
			wait := time.Minute * model.Config.Cron.Interval
			log.Info(`[Crontab] Waiting `, wait.String())
			time.Sleep(wait)
		}
	}()

}

func authCheck(h echo.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		if user, _ := c.Session().Get(`user`).(string); len(user) > 0 {
			c.Set(`user`, user)
			return h.Handle(c)
		}

		//临时认证
		if err := login(c, false); err != nil {
			return errors.New(`请登录`)
		}
		return h.Handle(c)
	}
}

func login(c echo.Context, saveSession bool) error {
	user := c.Form(`user`)
	pass := c.Form(`pass`)
	if pwd, ok := model.Config.Sys.Accounts[user]; ok && pwd == pass {
		if saveSession {
			c.Session().Set(`user`, user)
		}
		return nil
	}
	return errors.New(`登录失败，用户名或密码不正确`)
}

func addRouter(e *echo.Echo) {
	e.Get(`/`, func(c echo.Context) error {
		return c.String(`OK`)
	})
	e.Get(`/notice/list`, NoticeList, authCheck)
	e.Post(`/notice/add`, NoticeAdd, authCheck)
	e.Post(`/notice/edit`, NoticeEdit, authCheck)
	e.Get(`/notice/del`, NoticeDel, authCheck)

	e.Get(`/event/list`, EventList, authCheck)
	e.Post(`/event/add`, EventAdd, authCheck)
	e.Post(`/event/edit`, EventEdit, authCheck)
	e.Get(`/event/del`, EventDel, authCheck)

	e.Get(`/client/list`, ClientList, authCheck)
	e.Post(`/client/add`, ClientAdd, authCheck)
	e.Post(`/client/edit`, ClientEdit, authCheck)
	e.Get(`/client/del`, ClientDel, authCheck)

	e.Get(`/login`, func(c echo.Context) error {
		return JSON(c, nil, login(c, true))
	})
}
