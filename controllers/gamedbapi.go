package controllers

import (
	"github.com/astaxie/beego"
	// "encoding/json"
	"github.com/ro4tub/gamedb/models"
	. "github.com/ro4tub/gamedb/util"
)


type GamedbapiController struct {
	beego.Controller
}

func init()  {
	Log.Info("dburl: %s", beego.AppConfig.String("dburl"))
}

// /search?v=VALUE&p=PLATFORM&g=GENRE&t=TAG
func (this *GamedbapiController) Search()  {
	value := this.GetString("v")
	platform := this.GetString("p")
	genre := this.GetString("g")
	tag := this.GetString("t")
	Log.Info("search: value=%s, platform=%s, genre=%s, tag=%s\n", value, platform, genre, tag)
	if value == "" {
		this.Data["json"] = "invalid arguments"
		this.ServeJson()
		return
	}
	
	// this.Ctx.WriteString("hello")
	
	err, g := models.GetGamesByKeyword(value, platform, genre, tag)
	if err != nil {
		Log.Error("GetGamesByKeyword failed: %v", err)
		this.Data["json"] = err
	} else {
		this.Data["json"] = g
	}
	this.ServeJson()
}

// /game/id/
func (this *GamedbapiController) GetGameDetail() {
	id := this.GetString(":id")
	Log.Info("game: id=%v\n", id)
	// this.Ctx.WriteString("world")
	err, g := models.GetGameByUrlName(id)
	if err != nil {
		Log.Error("GetGameByUrlName failed: %v", err)
		this.Data["json"] = err
	} else {
		this.Data["json"] = g
	}
	this.ServeJson()
}


