package controllers

import (
	"github.com/astaxie/beego"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	oauthModel "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	//oredis "github.com/go-oauth2/redis/v4"
	//"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
)

type BaseController struct {
	beego.Controller
	ClientID string
	Scope    string
}

func (c *BaseController) Success() {
	var r response
	r.Code = 0
	r.Message = "success"
	c.Data["json"] = r
	c.ServeJSON()
	return
}

func (c *BaseController) Failed(code int, message string) {
	var r response
	r.Code = code
	r.Message = message
	c.Data["json"] = r
	c.ServeJSON()
	return
}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	Manager     = ManagerInit()
	ClientStore = ClientStoreInit()
	Srv         = SrvInit()
)

func ManagerInit() *manage.Manager {
	//生成oauth2 manager
	Manager := manage.NewDefaultManager()

	//配置config， 超时时间以及刷新时间
	Manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	return Manager
}

func ClientStoreInit() *store.ClientStore {
	// 生成token memory store

	//内存存储
	Manager.MustTokenStorage(store.NewMemoryTokenStore())

	//redis存储
	//Manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
	//	Addr: "127.0.0.1:6379",
	//	DB:   15,
	//}))

	// 生成client memory store
	ClientStore := store.NewClientStore()
	return ClientStore
}

func SrvInit() *server.Server {
	// 注册clientStorage
	Manager.MapClientStorage(ClientStore)

	// 生成server
	Srv := server.NewDefaultServer(Manager)

	//配置允许对令牌的GET请求
	Srv.SetAllowGetAccessRequest(true)

	// 配置允许从请求获取客户端
	Srv.SetClientInfoHandler(server.ClientFormHandler)

	// 配置刷新令牌config
	Manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	// 设置内部错误处理函数
	Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	// 设置响应错误处理函数
	Srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	return Srv
}

func (c *BaseController) Token() {
	if err := Srv.HandleTokenRequest(c.Ctx.ResponseWriter, c.Ctx.Request); err != nil {
		c.Failed(100001, err.Error())
		return
	}
}

func (c *BaseController) Credentials() {
	clientId := uuid.New().String()[:8]
	clientSecret := uuid.New().String()[:8]
	err := ClientStore.Set(clientId, &oauthModel.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://127.0.0.1:8080",
	})
	if err != nil {
		c.Failed(10001, err.Error())
		return
	}

	c.Data["json"] = struct {
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	c.ServeJSON()
	return
}

func (c *BaseController) validateToken() (oauth2.TokenInfo, error) {
	return Srv.ValidationBearerToken(c.Ctx.Request)
}

func (c *BaseController) Prepare() {
	if controller, _ := c.GetControllerAndAction(); controller == "BaseController" {
		return
	}
	info, err := c.validateToken()
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.WriteString(err.Error())
		return
	}
	c.ClientID = info.GetClientID()
	c.Scope = info.GetScope()
}
