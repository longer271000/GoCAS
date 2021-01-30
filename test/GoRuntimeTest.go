package test

import (
	"fmt"
	"runtime"
	"sync"
)
func smain()  {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=0;i<100;i++{
			fmt.Printf("线程1=====%d\n",i)
		}
	}()
	go func() {
		defer  wg.Done()
		for i:=0;i<100;i++{
			fmt.Printf("线程2=====%d\n",i)
		}
	}()
	wg.Wait()
}
