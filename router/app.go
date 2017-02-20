package router

import (
	"github.com/astaxie/beego"
        "API-SERVER/controller"
)

func init(){

	beego.Router("v1/cae/create", &controller.Appcontroller{})

	beego.Router("v2/apps/:appguid/bits", &controller.Appcontroller{} ,"Post:Uploadbits")

	beego.Router("v2/apps/:appguid/start", &controller.Appcontroller{} , "Post:Appstart")

	beego.Router("v2/apps/stage", &controller.Appcontroller{} , "Post:Stage")

	beego.Router("v2/apps", &controller.Appcontroller{} , "Get:GetAll")
}
