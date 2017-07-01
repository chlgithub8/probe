package main

import (
	"fmt"
	_ "probe/models"
	_ "probe/routers"
	"server_rc/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	dbServer := "127.0.0.1:3306"
	orm.RegisterDataBase(models.Db_Account, "mysql", fmt.Sprintf("wanchuanglab:czwl1812@tcp(%s)/account?charset=utf8", dbServer))
	orm.RegisterDataBase("probe", "mysql", fmt.Sprintf("probe:12345678@tcp(%s)/probe?charset=utf8", dbServer))
	//orm.RunSyncdb("probe", false, true)
}
func main() {
	beego.Run()

}
