package model

import (
	"time"

	"github.com/webx-top/db"
	"github.com/webx-top/echo"

	"github.com/webx-top/echodb_api_demo/app/dbschema"
)

func NewNotice(c echo.Context) *Notice {
	return &Notice{
		Notice:  &dbschema.Notice{},
		context: c,
	}
}

type Notice struct {
	*dbschema.Notice
	context echo.Context
}

type NoticeData struct {
	*dbschema.Notice
	Client *dbschema.Client `json:"client" xml:"client"`
	Event  *dbschema.Event  `json:"event" xml:"event"`
}

func (n *Notice) ListData(mw func(db.Result) db.Result, page, size int, args ...interface{}) ([]*NoticeData, func() int64, error) {
	r := []*NoticeData{}
	cnt, err := n.Notice.Param().SetCache(time.Minute*5).Model().List(nil, mw, page, size, args...)
	if err != nil {
		return r, cnt, err
	}
	ns := n.Notice.Objects()
	clientIds := []interface{}{}
	eventIds := []interface{}{}
	clientIdk := map[uint][]int{}
	eventIdk := map[uint64][]int{}
	r = make([]*NoticeData, len(ns))
	for k, v := range ns {
		if v.EventId > 0 {
			if _, ok := eventIdk[v.EventId]; !ok {
				eventIdk[v.EventId] = []int{}
				eventIds = append(eventIds, v.EventId)
			}
			eventIdk[v.EventId] = append(eventIdk[v.EventId], k)
		}
		if v.ClientId > 0 {
			if _, ok := clientIdk[v.ClientId]; !ok {
				clientIdk[v.ClientId] = []int{}
				clientIds = append(clientIds, v.ClientId)
			}
			clientIdk[v.ClientId] = append(clientIdk[v.ClientId], k)
		}
		r[k] = &NoticeData{
			Notice: v,
		}
	}
	if len(eventIds) > 0 {
		event := &dbschema.Event{}
		_, err = event.Param().SetCache(time.Minute*5).Model().List(nil, nil, 1, len(eventIds), db.Cond{"id IN": eventIds})
		if err != nil {
			return r, cnt, err
		}
		for _, v := range event.Objects() {
			if keys, ok := eventIdk[v.Id]; ok {
				for _, k := range keys {
					r[k].Event = v
				}
			}
		}
	}
	if len(clientIds) > 0 {
		client := &dbschema.Client{}
		_, err = client.Param().SetCache(time.Minute*5).Model().List(nil, nil, 1, len(clientIds), db.Cond{"id IN": clientIds})
		if err != nil {
			return r, cnt, err
		}
		for _, v := range client.Objects() {
			if keys, ok := clientIdk[v.Id]; ok {
				for _, k := range keys {
					r[k].Client = v
				}
			}
		}
	}
	return r, cnt, err
}
