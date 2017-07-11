package probeModels

import (
	"github.com/astaxie/beego/orm"
)

type Car struct {
	Id    int
	Key   string
	Value string
	Type  string
	Rate  float32
}

func init() {
	orm.RegisterModel(new(Car))
}
