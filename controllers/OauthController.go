package controllers

type ResourceController struct {
	BaseController
}

func (c *ResourceController) Get() {
	c.TplName = "index.html"
}
