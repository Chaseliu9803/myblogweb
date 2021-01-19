package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：chasel_6"
	c.Data["qq"] = "QQ：867966307"
	c.Data["tel"] = "Tel：1562616xxxx"
	c.TplName = "aboultme.html"
}
