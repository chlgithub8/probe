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
	inStr, _ := c.GetInt32("idg", 0)
	o := orm.NewOrm()
	o.Using("probe")
	qs := o.QueryTable("car")

	var cars []probeModels.Car
	totalcount, _ := qs.Filter("id__gt", 1).Count()
	count, err := qs.Filter("id__gt", inStr).Limit(-1).OrderBy("type", "id").All(&cars)
	if err == nil {
		result := make(map[string]interface{})
		result["cars"] = cars
		jsons, _ := json.Marshal(result)
		str := "totalcount:" + strconv.Itoa(int(totalcount)) + " show count:" + strconv.Itoa(int(count)) + "\n"
		c.Ctx.WriteString(str + string(jsons))
		c.Ctx.WriteString("\n\n\n\n\n\n")
		resultCar := make(map[string]interface{})
		var typeRank string
		var jsonCar []byte
		for _, v := range cars {
			resultCar["id"] = v.Id
			resultCar["key"] = v.Key
			resultCar["value"] = v.Value
			resultCar["type"] = v.Type
			resultCar["rate"] = v.Rate
			jsonCar, _ = json.Marshal(resultCar)
			typeRank += string(jsonCar)
			typeRank += ",\n"

		}
		c.Ctx.WriteString(typeRank)
	}

}
