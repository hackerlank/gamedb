package controllers

import (
	"os/exec"
	"bytes"
	"github.com/astaxie/beego"
	. "github.com/ro4tub/gamedb/util"
)

type AdminController struct {
	beego.Controller
}

type Result struct {
	Ret	string
	Desc string
}

// /admin/publish
func (this *AdminController) Publish()  {
	Log.Info("admin publish")
	cmd := exec.Command("bash", "deploy_gamedb.sh")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		this.Data["json"] = Result{Ret:"Failed", Desc:err.Error()}
		Log.Error("publish failed: %v", err)
	} else {
		this.Data["json"] = Result{Ret:"OK", Desc:out.String()}
		Log.Info("publish result: %s", out.String())
	}
	this.ServeJson()
}


