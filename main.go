package main

import (
	"flag"
	"fmt"

	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/standard"
	"github.com/webx-top/echo/middleware"
	"github.com/webx-top/echo/middleware/session"

	"github.com/webx-top/echodb_api_demo/app"
	"github.com/webx-top/echodb_api_demo/app/model"
)

func main() {
	model.CmdConfig.Port = flag.Int(`p`, 8080, `端口`)
	model.CmdConfig.Conf = flag.String(`c`, `conf/config.yaml`, `配置文件路径`)
	flag.Parse()

	e := echo.New()
	sessionOptions := &echo.SessionOptions{
		Name:   "SID",
		Engine: "cookie",
		CookieOptions: &echo.CookieOptions{
			Path:     "/",
			HttpOnly: false,
		},
	}
	e.Use(middleware.Log(), middleware.Recover(), session.Middleware(sessionOptions))
	app.Initialize(e)
	e.Run(standard.New(fmt.Sprintf(`:%v`, *model.CmdConfig.Port)))
}
