package routers

import (
	"oauth2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})

	//用于获取资源，当没有access_token时，无法获取
	beego.Router("/", &controllers.OauthController{})

	//用于通过clientID和clientSecret 获取token ， 通过token才能获取资源
	beego.Router("/token", &controllers.MainController{}, "get:Token")

	//用于生成clientID和clientSecret
	beego.Router("/credentials", &controllers.MainController{}, "get:Credentials")
}
