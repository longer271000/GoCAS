package datasource

import (
	"GoCAS/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)
/****
实例化数据库引擎方法：mysql 的数据引擎
****/

func NewMysqlEngine() *xorm.Engine {
	//数据库引擎
	engine, err := xorm.NewEngine("mysql","root:123@/iris?charset=utf8")

	//根据实体创建表
	err = engine.Sync2(
		new(models.LoginInfo),
		)

	if err != nil {
		panic(err.Error())
	}
	//设置显示SQL语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)
}