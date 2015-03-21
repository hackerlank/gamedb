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

type SearchResult struct {
	Ret 	string
	Games 	[]models.Game
	PlatformOutline map[string]int
	GenreOutline map[string]int
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
	result := SearchResult{}
	err, g := models.GetGamesByKeyword(value, platform, genre, tag)
	if err != nil {
		Log.Error("GetGamesByKeyword failed: %v", err)
		result.Ret = err.Error()
	} else {
		result.Ret = "ok"
		result.PlatformOutline = make(map[string]int)
		result.GenreOutline = make(map[string]int)
		Log.Warn("count: %d", len(g))
		if len(g) < 30 {
			result.Games = g
		} else {
			for i := 0; i < len(g); i++ {
				if i < 30 {
					result.Games = append(result.Games, g[i])
					continue
				}
				platform := g[i].Platform
				genre := g[i].Genre
				count, ok := result.PlatformOutline[platform]
				if !ok {
					result.PlatformOutline[platform] = 1
				} else {
					result.PlatformOutline[platform] = count + 1
				}
				count, ok = result.GenreOutline[genre]
				if !ok {
					result.GenreOutline[genre] = 1
				} else {
					result.GenreOutline[genre] = count + 1
				}
			}
		}
	}
	this.Data["json"] = result
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


