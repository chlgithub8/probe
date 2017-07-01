package controllers

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"probe/models"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) GetJson() {
	cstr, _ := c.GetInt32("idg", 0)
	o := orm.NewOrm()
	o.Using("probe")
	qs := o.QueryTable("car")
	var cars []model_probe.Car
	totalcount, _ := qs.Filter("id__gt", 1).Count()
	count, err := qs.Filter("id__gt", cstr).All(&cars)
	// countstr := strconv.Itoa(int(count))
	// c.Ctx.Output.Body([]byte(countstr))
	if err == nil {
		result := make(map[string]interface{})
		result["cars"] = cars
		json, _ := json.Marshal(result)
		str := "totalcount:" + strconv.Itoa(int(totalcount)) + " show count:" + strconv.Itoa(int(count)) + "\n"
		c.Ctx.WriteString(str + string(json))
	}

}
