//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
)

type Client struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*Client
	
	Id           	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Name         	string  	`db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	RecvUrl      	string  	`db:"recv_url" bson:"recv_url" comment:"接收通知的网址（支持占位符：{type},{event},{id}）" json:"recv_url" xml:"recv_url"`
	TypeMagazine 	uint    	`db:"type_magazine" bson:"type_magazine" comment:"是否(1/0)接收杂志通知" json:"type_magazine" xml:"type_magazine"`
	TypeBook     	uint    	`db:"type_book" bson:"type_book" comment:"是否(1/0)接收图书通知" json:"type_book" xml:"type_book"`
	TypeArticle  	uint    	`db:"type_article" bson:"type_article" comment:"是否(1/0)接收文章通知" json:"type_article" xml:"type_article"`
	EventAdd     	uint    	`db:"event_add" bson:"event_add" comment:"是否(1/0)接收新增通知" json:"event_add" xml:"event_add"`
	EventEdit    	uint    	`db:"event_edit" bson:"event_edit" comment:"是否(1/0)接收修改通知" json:"event_edit" xml:"event_edit"`
	EventDel     	uint    	`db:"event_del" bson:"event_del" comment:"是否(1/0)接收删除通知" json:"event_del" xml:"event_del"`
	PreHook      	string  	`db:"pre_hook" bson:"pre_hook" comment:"通知之前的处理钩子" json:"pre_hook" xml:"pre_hook"`
	HookParam    	string  	`db:"hook_param" bson:"hook_param" comment:"钩子参数" json:"hook_param" xml:"hook_param"`
	Disabled     	uint    	`db:"disabled" bson:"disabled" comment:"禁用" json:"disabled" xml:"disabled"`
}

func (this *Client) Trans() *factory.Transaction {
	return this.trans
}

func (this *Client) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *Client) Objects() []*Client {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *Client) NewObjects() *[]*Client {
	this.objects = []*Client{}
	return &this.objects
}

func (this *Client) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("client").SetModel(this)
}

func (this *Client) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *Client) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *Client) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *Client) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Client) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *Client) Add() (pk interface{}, err error) {
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		}
	}
	return
}

func (this *Client) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Update()
}

func (this *Client) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	},func(){
		this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		}
	}
	return 
}

func (this *Client) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetMiddleware(mw).Delete()
}

