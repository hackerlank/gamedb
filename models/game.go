package models;

import (
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

type Game struct {
	Id	int
	Name	string
	UrlName	string
	Pic		string
	SimpleDesc	string
	DetailDesc  string
	Platform	string // 机种
	Genre 		string // 游戏类型
	ReleaseDate	string // 发布日期
	tags	string
}

var (
	myorm orm.Ormer = nil
)

func init() {
	orm.RegisterModel(new(Game))
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dburl"), 30)
	myorm = orm.NewOrm()
}


func GetGameByUrlName(urlname string) (error, Game) {
	var game Game
	err := myorm.Raw("select * from game where url_name = ?", urlname).QueryRow(&game)
	return err, game
}

func GetGamesByKeyword(word string, platform string,  genre string, tag string) (error, []Game) {
	var games []Game
	qs := myorm.QueryTable("game")
	cond := orm.NewCondition()
	var cond1, cond2, cond3, cond4 *orm.Condition
	if platform != "" {
		cond1 = cond.And("platform__exact", platform)
	}
	if genre != "" {
		cond2 = cond.And("genre__exact", genre)
	}
	if tag != "" {
		cond3 = cond.And("tags__icontains", tag)
	}
	
	cond4 = cond.And("simple_desc__icontains", word).Or("name__icontains", word)
	
	qs.SetCond(cond.AndCond(cond1).AndCond(cond2).AndCond(cond3).AndCond(cond4)).All(&games)
	return nil, games
}