package test

import (
	"fmt"
)

func test(c chan int)  {

	c<-2

}
func write(c chan  int)  {
	x:=<-c
	fmt.Printf("管道数据====》%d",x)
}

func main2()  {
	ch:=make(chan int)

	go test(ch)
	go write(ch)



}
