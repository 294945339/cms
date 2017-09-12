package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"cms/config"
	"fmt"
	"log"
	"strings"
)

const BatchSize int = 500

var engine *xorm.Engine

func init() {
	conf := config.AppConfig.MySQL
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Url, conf.Database)
	e, err := xorm.NewEngine("mysql", dataSourceName)
	e.SetMapper(LowerFirstMapper{})
	e.SetMaxIdleConns(conf.MaxIdle)
	e.SetMaxOpenConns(conf.MaxActive)

	if config.AppConfig.Server.LogModelEnable {
		e.ShowSQL(true)
	} else {
		e.ShowSQL(false)
	}
	if err != nil {
		log.Fatalf("mysql connection failed: %q", err)
	}
	//if config.AppConfig.Server.LogModelEnable {
	//	engine.Logger().SetLevel(core.LOG_DEBUG)
	//} else {
	//	engine.Logger().SetLevel(core.LOG_ERR)
	//}
	engine = e
}

type LowerFirstMapper struct {
}

func (m LowerFirstMapper) Obj2Table(o string) string {
	return strings.ToLower(o[:1]) + o[1:]
}

func (m LowerFirstMapper) Table2Obj(t string) string {
	return t
}