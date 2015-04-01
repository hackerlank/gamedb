package routers

import (
	"github.com/ro4tub/gamedb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	
	beego.Router("/gamedbapi/search", &controllers.GamedbapiController{}, "get:Search")
	beego.Router("/gamedbapi/game/:id", &controllers.GamedbapiController{}, "get:GetGameDetail")
	beego.Router("/gamedbapi/game", &controllers.GamedbapiController{}, "post:SaveGameDetail")
	beego.Router("/admin/publish", &controllers.AdminController{}, "get:Publish")
}
