package controller

import (
	"silky-log/bean"

	"fmt"
	"silky-log/schema"

	iris "gopkg.in/kataras/iris.v6"
)

type LogCtrl struct {
}

var logBean = &bean.LogBean{}

func (logCtrl *LogCtrl) GetAll(ctx *iris.Context) {
	if list, err := logBean.FindAll(); err != nil {
		ctx.SetStatusCode(500)
		ctx.Writef("查询数据失败")
		return
	} else {
		ctx.JSON(200, list)
	}
}

func (logCtrl *LogCtrl) Post(ctx *iris.Context) {

	data := &schema.Log{}
	if err := ctx.ReadJSON(data); err != nil {
		ctx.SetStatusCode(500)
		fmt.Println(err)
		ctx.Done()
		return
	}
	data.IP = ctx.RemoteAddr()
	logBean.Add(data)
	ctx.SetStatusCode(200)
	ctx.Done()
	// if vps, err := logBean.Get(id); err != nil {
	// 	ctx.SetStatusCode(403)
	// 	ctx.Writef("数据不存在")
	// } else {
	// 	ctx.JSON(200, vps)
	// }
}
