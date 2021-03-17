package controllers

type ResourceController struct {
	BaseController
}

func (c *ResourceController) Get() {
	c.TplName = "index.html"
}

func (c *ResourceController) Auth() {
	if _, err := c.validateToken(); err != nil {
		c.Failed(100001, err.Error())
		return
	}
	c.Success(c.TokenInfo)
}
