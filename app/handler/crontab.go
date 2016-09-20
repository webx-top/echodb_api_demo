package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/admpub/log"
	"github.com/webx-top/db"
	"github.com/webx-top/echodb_api_demo/app/dbschema"
)

func Crontab() (complete bool) {
	notice := &dbschema.Notice{}
	client := &dbschema.Client{}
	event := &dbschema.Event{}
	page := 1
	size := 500
	pages := 0
	fn := func(r db.Result) db.Result {
		return r.OrderBy(`id`)
	}
	cnt, err := event.List(nil, fn, page, size, db.Cond{"finished": 0})
	if err != nil {
		log.Error(err)
	}
	events := event.Objects()
	if len(events) > 0 && pages == 0 {
		pages = Pages(cnt(), size)
	}

	for ; page <= pages; page++ {
		if page > 1 {
			_, err = event.List(nil, fn, page, size, db.Cond{"finished": 0})
			if err != nil {
				log.Error(err)
			}
			events = event.Objects()
		}
		for _, evt := range events {
			cond := db.Cond{"disabled": 0, "recv_url !=": ""}
			switch evt.Name {
			case "add", "edit", "del":
				cond["event_"+evt.Name] = 1
			default:
				log.Warn(`无效的事件名：`, evt.Name)
				continue
			}
			switch evt.Target {
			case "magazine", "book", "article":
				cond["type_"+evt.Target] = 1
			default:
				log.Warn(`无效的Target名：`, evt.Target)
				continue
			}
			_, err = client.List(nil, nil, 1, 9999999, cond)
			if err != nil {
				log.Error(err)
				continue
			}
			clients := client.Objects()
			if len(clients) == 0 {
				evt.Finished = uint(time.Now().Unix())
				err = evt.Edit(nil, db.Cond{"id": evt.Id})
				if err != nil {
					log.Error(err)
				}
				continue
			}
			wg := &sync.WaitGroup{}
			wg.Add(len(clients))
			for _, c := range clients {
				go sendNotice(c, notice, evt, wg)
			}
			wg.Wait()
			total, err := notice.Param().SetArgs(db.Cond{"event_id": evt.Id, "finished": 0}).Count()
			if err != nil {
				log.Error(err)
			} else if total == 0 {
				evt.Finished = uint(time.Now().Unix())
				err = evt.Edit(nil, db.Cond{"id": evt.Id})
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
	complete = true
	return
}

func sendNotice(c *dbschema.Client, notice *dbschema.Notice, evt *dbschema.Event, wg *sync.WaitGroup) error {
	defer wg.Done()
	valid := true
	if len(c.PreHook) > 0 {
		switch c.PreHook {
		case `checkValid`:
			valid = checkValid(strings.Split(c.HookParam, `,`)...)
		default:
			log.Error(errors.New(`无效的hook名：` + c.PreHook))
			return nil
		}
	}
	if !valid {
		return nil
	}
	ncond := db.Cond{"client_id": c.Id, "event_id": evt.Id}
	err := notice.Get(nil, ncond)
	if err != nil {
		if err == db.ErrNoMoreRows {
			notice = &dbschema.Notice{
				ClientId: c.Id,
				EventId:  evt.Id,
			}
			_, err = notice.Add()
		}
		if err != nil {
			log.Error(err)
		}
	}
	if notice.Finished > 0 {
		return err
	}
	notice.Retry++
	noticeURL := strings.Replace(c.RecvUrl, "{type}", evt.Target, -1)
	noticeURL = strings.Replace(noticeURL, "{event}", evt.Name, -1)
	noticeURL = strings.Replace(noticeURL, "{id}", fmt.Sprintf(`%d`, evt.TargetId), -1)
	var resp *http.Response
	resp, err = http.Get(noticeURL)
	if err == nil {
		if resp.StatusCode != http.StatusOK {
			err = errors.New(resp.Status + `：` + noticeURL)
		}
	}
	if err != nil {
		log.Error(err)
	} else {
		log.Info(`Successfully sent a notification to `, noticeURL, ` [`, evt.Name, `:`, evt.Target, `:`, evt.TargetId, `]`)
		notice.Finished = uint(time.Now().Unix())
	}
	err = notice.Edit(nil, ncond)
	if err != nil {
		log.Error(err)
	}
	return err
}

func checkValid(args ...string) bool {
	return true
}
