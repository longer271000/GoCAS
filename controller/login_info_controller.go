package controller

import (
	"GoCAS/service"
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type LoginInfoController struct {
	//iris框架 自动为每一个请求绑定上下文对象，可以做接受参数
	Ctx iris.Context
	//功能实体,引入Service接口
	Service service.LoginInfoService

	//session对象；存储session信息
	Session *sessions.Session



}
const (
	LoginKey = "LoginKey"
)
type LoginInfo struct {
	LoginName string "json:'loginName'"
	LoginPwd string "json:'loginPwd'"
}
/****
json 请求格式
接口/login
***/
func (ac *LoginInfoController) PostLogin(context iris.Context) mvc.Result {
	var loginInfo LoginInfo
	ac.Ctx.ReadJSON(&loginInfo)

	//数据校验
	if loginInfo.LoginName == "" || loginInfo.LoginPwd == ""{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":"0",
				"success":"登录失败",
				"message":"用户名或者密码为空",
			},
		}
		//根据用户名，密码查询数据库
		loginInfoData,exist :=ac.Service.GetByLoginNamePwd(loginInfo.LoginName,loginInfo.LoginPwd)

		//用户是否存在
		if !exist{
			return mvc.Response{
				Object: map[string]interface{}{
					"status":"1",
					"success":"登录失败",
					"message":"用户名或者密码错误！",
				},
			}
		}
		//用户名和密码正确
		userByte,_ := json.Marshal(loginInfoData)
		ac.Session.Set(LoginKey,userByte)

		return mvc.Response{
			Object: map[string]interface{}{
				"status":"1",
				"success":"登录成功",
				"message":"登录成功",
			},
		}
	}

}
