package test

import (
	"fmt"
	"strconv"
)

func ssmain()  {
	fmt.Println("this first go run file")
	//var array [5]intarray = [5]int{1,2,3,4,5}
	//array := [5]int{1: 1, 3: 4}
	//for i, v := range array {
	//	fmt.Printf("索引:%d,值:%d\n", i, v)
	//}
	//var array [5]intList = [5]int{1,2,3,4,5}
	//数组
	//array := [5]int{1,2,3,4,5}
	//for i :=0;i<5;i++{
	//	fmt.Printf("索引:%d;值:%d\n",i,array[i])
	//}
	//
	//切片
	//silList := []int{1,2,3,4,5}
	//for i :=0;i< len(silList);i++{
	//	fmt.Printf("切片数据%d\n",silList[i])
	//}
	//silList = append(silList, 6)
	//for i :=0;i< len(silList);i++{
	//	fmt.Printf("增加切片数据%d\n",silList[i])
	//}
	//Map
	hashMap:=make(map[string]int)
	hashMap["张1"]=41
	hashMap["张2"]=42
	hashMap["张3"]=43
	hashMap["张4"]=44

	var arg=hashMap["张3"]
	fmt.Printf("哈希值==%d\n",arg)
	delete(hashMap,"张3")
	//age,exists := dict["李四"]
	//arg,exist:=hashMap["张3"]
	//fmt.Printf("哈希值==%d",arg)

	//dict :=map[string]int{"a":1,"b":2}
//函数，方法
 //name :="朱华"
 //fmt.Println(modify(name))
 sum:= add(1,3)
 fmt.Println(sum)
//结构
 p:= person{name: "Jack",sex: 1,age: 0}
 fmt.Println(p.String())
 p.modifyPerson()
	fmt.Println(p.String())

 sum2,ex:= addException(1,2)
 fmt.Printf("双参数==%d===>%s",sum2,ex)

	print(1,2,3,4)
	print(1,2,3,4,5,6)

 //接口实现
 var c cat
 //c:=1000
 invoke(c)
 //嵌入式
  userinfo:= admin{user{"朱华",1},"admin"}

	UerInfoInvoke(userinfo)
  //

 


}
func add(a int,b int) int  {
	return a+b

}
func modify(s string) string {
	s=s+s
	return s
}
type person struct {
	name string
	sex int
	age int
}

func (p person)String() string {
	return "name="+p.name+";sex="+strconv.Itoa( p.sex)+";age="+strconv.Itoa(p.age)
	
}
/**
指针
**/
func (p *person) modifyPerson()  {
	p.name="李思思"
}
//多参数返回
func addException(a,b int)(int,error)  {
	return a+b,nil

}
//可变参数
func print(s ...interface{})  {
	for _,v :=range s{
		fmt.Printf("可变参数===%d\n",v)
	}
}
//接口实现
type cat int

type animal interface {
	printinfo()
}

func invoke(a animal) {

	a.printinfo()
}
func (a cat) printinfo() {
	fmt.Println("接口实现打印：this a cat")
}
//嵌入类型
type user struct {
	name string
	age int
}
type admin struct{
	userinfo user
	usertype string
}
type sayHello interface {
	printUserInfo()
}

func UerInfoInvoke(s sayHello)  {
	s.printUserInfo()
}
func (userinfo admin)printUserInfo()  {
	fmt.Printf("this userinfo user tpye admin,name=%s",userinfo.userinfo.name)
}