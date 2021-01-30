package test

import (
	"fmt"
	"reflect"
)

func main5()  {

	 userinfo:=User{name: "hua.zhu",age: 1}

	u:=reflect.ValueOf(&userinfo)
	//u.Elem().Set()

	mPrint:=u.MethodByName("PrintInfo")
	args:=[]reflect.Value{reflect.ValueOf("dddd")}

	mPrint.Call(args)

}
type User struct {
	name string
	age int
}

func (u User)PrintInfo(v string)  {
	fmt.Println("命名==="+u.name+"||||"+v)
}