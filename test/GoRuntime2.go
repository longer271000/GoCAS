package test

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg sync.WaitGroup
)
func main1()  {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Printf("GoRuntime2===count[%d]",count)

}
func incCount()  {
	defer wg.Done()
	for i:=0;i<2;i++{
		value:=atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		//count=value
		atomic.StoreInt32(&count,value)
	}

}

