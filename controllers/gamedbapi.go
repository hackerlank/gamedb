package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/ro4tub/gamedb/models"
	. "github.com/ro4tub/gamedb/util"
	"strings"
)


type GamedbapiController struct {
	beego.Controller
}

func init()  {
	Log.Info("dburl: %s", beego.AppConfig.String("dburl"))
}

type SearchOutline struct {
	Key	string
	Count int
}

type SearchResult struct {
	Ret 	string
	TotalCount int // 总个数
	Games 	[]models.Game
	PlatformOutline []SearchOutline
	GenreOutline []SearchOutline
	KeywordOutline []SearchOutline
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
		result.TotalCount = len(g)
		platformOutline := make(map[string]int)
		genreOutline := make(map[string]int)
		var nameMatched = 0
		
		Log.Warn("count: %d", len(g))
		for i := 0; i < len(g); i++ {
			if i < 30 {
				result.Games = append(result.Games, g[i])
			}
			if strings.Contains(g[i].Name, value) {
				nameMatched = nameMatched + 1
			}
			platform := g[i].Platform
			genre := g[i].Genre
			count, ok := platformOutline[platform]
			if !ok {
				platformOutline[platform] = 1
			} else {
				platformOutline[platform] = count + 1
			}
			count, ok = genreOutline[genre]
			if !ok {
				genreOutline[genre] = 1
			} else {
				genreOutline[genre] = count + 1
			}
		}
		for k, v := range platformOutline {
			result.PlatformOutline = append(result.PlatformOutline, SearchOutline{k, v})
		}
		for k, v := range genreOutline {
			result.GenreOutline = append(result.GenreOutline, SearchOutline{k, v})
		}
		result.KeywordOutline = append(result.KeywordOutline, SearchOutline{"name", nameMatched})
		result.KeywordOutline = append(result.KeywordOutline, SearchOutline{"simple_desc", result.TotalCount - nameMatched})
		
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

func (this *GamedbapiController) SaveGameDetail() {
	// id := this.GetString(":id")
	Log.Info("SaveGameDetail")
	g := models.Game{}
    if err := json.Unmarshal(this.Ctx.Input.RequestBody, &g); err != nil {
		Log.Error("json.Unmarshal failed: %v", err)
		this.Data["json"] = err
     } else {
		 Log.Info("save game info: %s", g.ToString())
		 models.SaveGame(g)
		 this.Data["json"] = g
     }
	this.ServeJson()
}



