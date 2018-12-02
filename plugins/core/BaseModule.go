package core

import (
	"github.com/ecdiy/goserver/webs"
	"github.com/ecdiy/goserver/utils"
	"github.com/ecdiy/goserver/gpa"
	"github.com/cihub/seelog"
	"github.com/ecdiy/goserver/plugins"
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
		plugins.ElementMap[m.MustAttr("Id")] = m
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

func (app *Module) Verify(ele *utils.Element) {
	vb := webs.NewVerify(ele, plugins.GetGpa(ele), putFunRun)
	put(ele, vb)
	tfn, ext := ele.AttrValue("TplFunName")
	if ext {
		webs.RegisterBaseFun(tfn, vb)
		seelog.Info("添加模版函数:", tfn)
	}
}
