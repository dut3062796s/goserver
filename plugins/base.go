package plugins

import (
	"github.com/ecdiy/goserver/utils"
	"github.com/ecdiy/goserver/gpa"
	"github.com/cihub/seelog"
)

type Module struct {
}

func (app *Module) Include(ele *utils.Element) {
	f := getFile(ele.Value)
	seelog.Info("include file:", f)
	dom, err := utils.LoadByFile(f)
	if err == nil {
		InvokeByXml(dom)
	} else {
		seelog.Error(err)
		panic("配置文件出错:" + f)
	}
}

func (app *Module) Map(ele *utils.Element) {
	ms := ele.AllNodes()
	for _, m := range ms {
		ElementMap[m.MustAttr("Id")] = m
	}
}

func (app *Module) Parameter(ele *utils.Element) {
	ps := ele.AllNodes()
	for _, p := range ps {
		utils.EnvParamSet(p.Name(), p.Value)
	}
}

func (app *Module) Gpa(ele *utils.Element) {
	dsn, b := ele.AttrValue("DbDsn")
	if b && len(dsn) > 0 {
		db := gpa.InitGpa(dsn)
		put(ele, db)
	}
}
