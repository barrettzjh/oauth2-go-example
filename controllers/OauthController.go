package controllers

type OauthController struct {
	MainController
}

func (c *OauthController) Get() {
	c.TplName = "index.tpl"
}