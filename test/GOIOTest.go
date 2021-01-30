package test

import (
	"bytes"
	"fmt"
	"os"
)


func main4()  {
	var b bytes.Buffer
	b.Write([]byte("你好，Go语言"))
	fmt.Fprintf(&b,",","http://wwww.baidu.com")
	b.WriteTo(os.Stdout)
}
