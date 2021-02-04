package service

import (
	"GoCAS/models"
	"fmt"
	"github.com/go-xorm/xorm"
)

type LoginInfoService interface {

	GetByLoginNamePwd(loginname,password string) (models.LoginInfo,bool)

}
type loginInfoService struct {
	engine *xorm.Engine
}


//实现接口
func NewLoginInfoService(db *xorm.Engine) LoginInfoService {
	return &loginInfoService{
		engine: db,
	}
	
}
/***
通过用户名和密码查询验证
***/
func (ac *loginInfoService)  GetByLoginNamePwd(loginname,password string)(models.LoginInfo,bool){
 var logininfo models.LoginInfo
 ac.engine.Where(" login_name = ? and login_pwd = ?",loginname,password).Get(&logininfo)
 fmt.Println(logininfo,".........",logininfo.ID != 0)
 return logininfo,logininfo.ID !=0
}